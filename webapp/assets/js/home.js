$('#nova-publicacao').on('submit', criarPost);
$('.curtir-postagem').on('click', curtirPost)

function criarPost(evento) {
    evento.preventDefault();
    $.ajax({
        url: '/publicacoes', // URL para onde enviar o POST
        type: 'POST', // Método HTTP a ser utilizado (GET, POST, etc.)
        dataType: 'json', // Tipo de dado esperado de retorno
        data: {
            titulo: $('#titulo').val(),
            conteudo: $('#conteudo').val(),
        }
    }).done(function(data) {
        alert("Publicação criada com sucesso");
        window.location.reload();
    }).fail(function(data) {
        console.log(data);
        alert("falha ao criada publicação");
    });
}

function curtirPost(evento) {
    evento.preventDefault();
    // Aqui vai recuperar o id da publicação
    const idPost = $(this).closest('.card').data('publicacao-id');

    $.ajax({
        url: `/publicacoes/${idPost}/curtir`, // URL para onde enviar o POST
        type: 'POST', // Método HTTP a ser utilizado (GET, POST, etc.)
        dataType: 'json', // Tipo de dado esperado de retorno
    }).done(function(data) {
        alert("Você curtiu essa Publicação");
        window.location.reload();
    }).fail(function(data) {
        console.log(data);
        alert("falha ao curtir publicação");
    });
}