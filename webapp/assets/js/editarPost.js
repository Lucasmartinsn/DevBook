$('#editar-publicacao').on('submit', editarPost);

function editarPost(evento) {
    evento.preventDefault();
    $(this).prop('disabled', true)
    const idPost = $(this).closest('.card').data('publicacao-id');

    $.ajax({
        url: `/publicacoes/${{idPost}}`, // URL para onde enviar o POST
        type: 'POST', // Método HTTP a ser utilizado (GET, POST, etc.)
        dataType: 'json', // Tipo de dado esperado de retorno
        data: {
            titulo: $('#titulo').val(),
            conteudo: $('#conteudo').val(),
        }
    }).done(function(data) {
        alert("Publicação atualizada com sucesso");
        window.location = '/perfil';
    }).fail(function(data) {
        console.log(data);
        alert("falha ao criada publicação");
    }).always(function() {
        $('#editar-publicacao').prop('disabled', false)
    }) ;
}