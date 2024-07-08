$('#criar-publicacao').on('submit', criarPost);

function criarPost(evento) {
    evento.preventDefault();
    console.log("criar publicação");
    // $.ajax({
    //     url: '/publicacoes', // URL para onde enviar o POST
    //     type: 'POST', // Método HTTP a ser utilizado (GET, POST, etc.)
    //     dataType: 'json', // Tipo de dado esperado de retorno
    //     data: {
    //         titulo: $('#titulo').val(),
    //         conteudo: $('#conteudo').val(),
    //     }
    // }).done(function(data) {
    //     alert("Publicação criada com sucesso");
    // }).fail(function(data) {
    //     console.log(data.responseJSON);
    //     alert("falha ao criada publicação");
    // });
}