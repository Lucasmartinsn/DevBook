$('#editar-publicacao').on('submit', editarPost);

function editarPost(evento) {
    evento.preventDefault();
    $(this).prop('disabled', true)
    var id = $('#postID').val();

    $.ajax({
        url: `/publicacoes/${id}`, // URL para onde enviar o POST
        type: 'PUT', // Método HTTP a ser utilizado (GET, POST, etc.)
        dataType: 'json', // Tipo de dado esperado de retorno
        data: {
            titulo: $('#titulo').val(),
            conteudo: $('#conteudo').val(),
        }
    }).done(function(data) {
        alert("Publicação atualizada com sucesso");
        history.back();
    }).fail(function(data) {
        console.log(data);
        alert("falha ao atualizar publicação");
    }).always(function() {
        $('#editar-publicacao').prop('disabled', false)
    }) ;
}