import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { MicroserviceOptions, Transport } from '@nestjs/microservices';
import { join } from 'path';

async function bootstrap() {
  const app = await NestFactory.createMicroservice<MicroserviceOptions>(AppModule, {
    transport: Transport.GRPC,
    options: {
      package: 'main',
      protoPath: join(__dirname, '../proto/order.proto'),
      url: 'localhost:6000',
    },
  });

  await app.listen();
  console.log('🚀 User Service is running on gRPC port 6000');
}
bootstrap();
