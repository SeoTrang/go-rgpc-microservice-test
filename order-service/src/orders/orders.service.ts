import { Inject, Injectable } from '@nestjs/common';
import { ClientGrpc } from '@nestjs/microservices';
import { lastValueFrom, Observable } from 'rxjs';

interface UserService {
  getUser(data: {
    id: number;
  }): Observable<{ id: number; name: string; age: number }>;
}

@Injectable()
export class OrdersService {
  private userService: UserService;
  private orders: { id: number; code: string; user_id: number }[] = [
    {
      id: 1,
      code: '#FTCD1',
      user_id: 1,
    },
    {
      id: 2,
      code: '#FTCD2',
      user_id: 1,
    },
    {
      id: 3,
      code: '#FTCD3',
      user_id: 2,
    },
    {
      id: 4,
      code: '#FTCD4',
      user_id: 1,
    },
    {
      id: 5,
      code: '#FTCD5',
      user_id: 2,
    },
    {
      id: 6,
      code: '#FTCD6',
      user_id: 3,
    },
  ];

  constructor(@Inject('USER_PACKAGE') private client: ClientGrpc) {}

  onModuleInit() {
    this.userService = this.client.getService<UserService>('UserService');
  }
  async getById(
    id: number,
  ): Promise<{ id: number; code: string; user_id: number; user?: any }> {
    const order = this.orders.find((item) => item.id === id);
    if (order) {
      const user: { id: number; name: string; age: number } =
        await lastValueFrom(this.userService.getUser({ id: order.user_id }));
      return {
        ...order,
        user,
      };
    }
    return order;
  }

  async getAll(): Promise<
    { id: number; code: string; user_id: number; user?: any }[]
  > {
    const ordersWithUsers = await Promise.all(
      this.orders.map(async (order) => {
        const user = await lastValueFrom(
          this.userService.getUser({ id: order.user_id })
        );
        // const user = {};
        return { ...order, user };
      }),
    );

    return ordersWithUsers;
  }
}
