$('#formulario-atualiza-cadastro').on('submit', atualizarUser);
$('#formulario-atualizar-senha-cadastro').on('submit', atualizarSenhaUser);

function atualizarUser(evento) {
    evento.preventDefault();
    var id = $('#userID').val();

    $.ajax({
        url: `/usuario/${id}`, // URL para onde enviar o POST
        type: 'PUT', // Método HTTP a ser utilizado (GET, POST, etc.)
        dataType: 'json', // Tipo de dado esperado de retorno
        data: {
            nome: $('#nome').val(),
            nick: $('#nick').val(),
            email: $('#email').val(),
        }
    }).done(function(data) {
        alert("cadastrado com sucesso");
        window.location.reload();
    }).fail(function(data) {
        console.log(data.responseJSON);
        alert("falha ao atualizar usuario: Nick ou email ja em uso");
    });
}

function atualizarSenhaUser(evento) {
    evento.preventDefault();
    var id = $('#userID').val();

    if ($('#senha').val() != $('#ConfirmarSenha').val()) {
        alert("senhas nao coencidem");
        return
    }

    $.ajax({
        url: `/usuario/${id}/atualizar-pass`, // URL para onde enviar o POST
        type: 'PUT', // Método HTTP a ser utilizado (GET, POST, etc.)
        dataType: 'json', // Tipo de dado esperado de retorno
        data: {
            atual: $('#atual').val(),
            senha: $('#senha').val(),
        }
    }).done(function(data) {
        alert("senha atualizada com sucesso");
        window.location.reload();
    }).fail(function(data) {
        console.log(data.responseJSON);
        alert("falha ao atualizar usuario");
    });
}