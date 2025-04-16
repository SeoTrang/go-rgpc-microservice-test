import { Controller } from '@nestjs/common';
import { GrpcMethod } from '@nestjs/microservices';
import { OrdersService } from './orders.service';

@Controller('orders')
export class OrdersController {
  constructor(private orderService: OrdersService) {}
  @GrpcMethod('OrderService', 'GetOrder')
  async getOrder(data: { id: string }): Promise<{
    id: number;
    code: string;
    user_id: number;
    user?: any;
  }> {
    return await this.orderService.getById(Number(data.id));
  }

  @GrpcMethod('OrderService', 'GetAllOrders')
  async getAll(): Promise<{
    orders: { id: number; code: string; user_id: number; user?: any }[];
  }> {
    const allOrder = await this.orderService.getAll();

    return { orders: allOrder };
  }
}
