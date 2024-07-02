$('#login').on('submit', fazerLogin);

function fazerLogin(evento) {
    evento.preventDefault();
    $.ajax({
        url: '/login', // URL para onde enviar o POST
        type: 'POST', // Método HTTP a ser utilizado (GET, POST, etc.)
        dataType: 'json', // Tipo de dado esperado de retorno
        data: {
            email: $('#email').val(),
            senha: $('#senha').val(),
            // adicione mais campos conforme necessário
        }
    }).done(function() {
        alert("Bem vindo ao Devbook");
        window.location = "/home"
    }).fail(function(data) {
        console.log(data.responseJSON);
        alert("falha ao fazer Login");
    });
}