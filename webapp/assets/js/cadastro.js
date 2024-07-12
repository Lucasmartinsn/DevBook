$('#formulario-cadastro').on('submit', criarUser);

function criarUser(evento) {
    evento.preventDefault();
    if ($('#senha').val() != $('#ConfirmarSenha').val()) {
        Swal.fire({
            title: "Senhas não Coincidem!",
            text: "Click no botão para tentar novamente!",
            icon: "error"
          });
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
        }
    }).done(function(data) {
        Swal.fire({
            title: "Cadastro realizado!",
            text: "Click no botão para voltar para a pagina de login!",
            icon: "success"
          }).then(() => {
            window.location = "/";
          });
    }).fail(function(data) {
        console.log(data.responseJSON);
        Swal.fire({
            title: "Falha ao realizar o Cadastro!",
            text: "Click no botão para tentar novamente!",
            icon: "error"
          });
    });

}