const amqp = require("amqplib");

async function consume() {
  try {
    const conn = await amqp.connect("amqp://localhost");
    const channel = await conn.createChannel();

    const queueName = "hello";
    channel.assertQueue(queueName, {
      durable: true,
    });
    channel.consume(
      queueName,
      channel.consume(
        queueName,
        function (msg) {
          console.log(" [x] Received %s", msg.content.toString());
          console.log(JSON.parse(`{"campo": 3, "coluna": "3trdrt"}`));
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
