$('#formulario-cadastro').on('submit', criarUser);

function criarUser(evento) {
    evento.preventDefault();
    if ($('#senha').val() != $('#ConfirmarSenha').val()) {
        alert("senhas nao coencidem");
        return
    }
    $.ajax({
        url: '/usuario', // URL para onde enviar o POST
        type: 'POST', // Método HTTP a ser utilizado (GET, POST, etc.)
        dataType: 'json', // Tipo de dado esperado de retorno
        data: {
            nome: $('#nome').val(),
            nick: $('#nick').val(),
            email: $('#email').val(),
            senha: $('#senha').val(),
            // adicione mais campos conforme necessário
        }
    }).done(function(data) {
        alert("usuario cadastrado com sucesso");
        window.location.reload();
        alert("Voltando para a pagina de Login");
        history.back();
    }).fail(function(data) {
        console.log(data.responseJSON);
        alert("falha ao cadastra usuario");
    });

}