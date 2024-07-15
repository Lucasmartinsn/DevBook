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
        Swal.fire({
            title: "Bem vindo ao Devbook!",
            icon: "success"
        }).then(() => {
            window.location = "/home"
        });
    }).fail(function(data) {
        console.log(data.responseJSON);
        Swal.fire({
            title: "Falha ao fazer Login!",
            icon: "error",
            text:"Usuario ou senha Invalidos!"
        });
    });
}