$('#formulario-cadastro').on('submit', criarUser);

function criarUser(evento) {
    evento.preventDefault();
    if ($('#senha').val() != $('#ConfirmarSenha').val()) {
        alert("senhas nao coencidem");
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
        },
        success: function (data) {
            // Função executada quando a requisição for bem-sucedida
            console.log('Requisição bem-sucedida:', data);
        },
        error: function (xhr, status, error) {
            // Função executada em caso de erro na requisição
            console.error('Erro na requisição:', status, error);
        }
    });

}