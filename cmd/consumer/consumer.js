const amqp = require("amqplib");

async function consume() {
  try {
    const conn = await amqp.connect("amqp://localhost");
    const channel = await conn.createChannel();

    const queueName = "hello";
    channel.assertQueue(queueName, {
      durable: false,
    });
    channel.consume(
      queueName,
      channel.consume(
        queueName,
        function (msg) {
          console.log(" [x] Received %s", msg.content.toString());
        },
        {
          noAck: true,
        }
      )
    );
  } catch (error) {
    console.log(error);
    return;
  }
}

consume();
