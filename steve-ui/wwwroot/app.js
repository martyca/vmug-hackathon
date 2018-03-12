(function () {
    var self = this;

    console.log("Hen.lo.");

    let connection = new signalR.HubConnection('/messagehub');

    connection.on('send', data => {
        console.log(data);
    });

    console.log(connection);
    connection.start().then(() => connection.invoke('send', 'Hello'));
})();