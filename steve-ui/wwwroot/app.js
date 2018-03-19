var vm = (function () {
    var self = this;
    self.connection = new signalR.HubConnection('/messagehub');

    self.messages = ko.observableArray();

    // Constructor

    self.connection.on("send", data => {
        self.messages.push(data);
    });

    self.connection.start() //.then(() => connection.invoke('send', 'Hello'));
})();

ko.applyBindings(vm, document.getElementById("app"));