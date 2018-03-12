using System;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.SignalR;
using steve_ui.Hubs;

namespace steve_ui.Controllers {

    [Route("api/messages/v1")]
    public class MessageController : Controller
    {
        private readonly IHubContext<MessageHub> _messageHub;

        public MessageController(IHubContext<MessageHub> messageHub)
        {
            _messageHub = messageHub;
        }

        [HttpPost, Route("add")]
        public async Task<ActionResult> AddMessage([FromForm] string message) {
            Console.WriteLine($"New message: { message }");

            await _messageHub.Clients.All.SendAsync("send", message);

            return await Task.FromResult(Ok());
        }
    }
}