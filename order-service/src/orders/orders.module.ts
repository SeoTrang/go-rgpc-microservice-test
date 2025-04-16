import { Module } from '@nestjs/common';
import { OrdersController } from './orders.controller';
import { OrdersService } from './orders.service';
import { ClientsModule, Transport } from '@nestjs/microservices';
import { join } from 'path';

@Module({
  imports: [
    ClientsModule.register([
      {
        name: 'USER_PACKAGE',
        transport: Transport.GRPC,
        options: {
          package: 'main',
          protoPath: join(__dirname, '../../proto/user.proto'),
          url: 'localhost:4000',
        },
      },
    ]),

  ],
  controllers: [OrdersController],
  providers: [OrdersService]
})
export class OrdersModule {}
