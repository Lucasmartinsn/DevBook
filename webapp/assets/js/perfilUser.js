$('#formulario-atualiza-cadastro').on('submit', atualizarUser);
$('#formulario-atualizar-senha-cadastro').on('submit', atualizarSenhaUser);

function atualizarUser(evento) {
    evento.preventDefault();
    $.ajax({
        url: '/usuario', // URL para onde enviar o POST
        type: 'POST', // Método HTTP a ser utilizado (GET, POST, etc.)
        dataType: 'json', // Tipo de dado esperado de retorno
        data: {
            nome: $('#nome').val(),
            nick: $('#nick').val(),
            email: $('#email').val(),
        }
    }).done(function(data) {
        alert("atualizado com sucesso");
        window.location.reload();
    }).fail(function(data) {
        console.log(data.responseJSON);
        alert("falha ao cadastra usuario");
    });
}

function atualizarSenhaUser(evento) {
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
            senha: $('#senha').val(),
        }
    }).done(function(data) {
        alert("senha atualizada com sucesso");
        window.location.reload();
    }).fail(function(data) {
        console.log(data.responseJSON);
        alert("falha ao cadastra usuario");
    });
}