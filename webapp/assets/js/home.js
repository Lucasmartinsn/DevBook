$('#nova-publicacao').on('submit', criarPost);
$('.curtir-post').on('click', curtirPost)
$('.descurtir-post').on('click', descurtirPost)

function criarPost(evento) {
    evento.preventDefault();
    Swal.fire({
        title: "Atenção!",
        text: "Vocer deseja realizar a Publicação?",
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "warning"
    }).then((confirmacao) => {
        if (!confirmacao.isConfirmed) return;
    
        $.ajax({
            url: '/publicacoes', // URL para onde enviar o POST
            type: 'POST', // Método HTTP a ser utilizado (GET, POST, etc.)
            dataType: 'json', // Tipo de dado esperado de retorno
            data: {
                titulo: $('#titulo').val(),
                conteudo: $('#conteudo').val(),
            }
        }).done(function(data) {
            Swal.fire({
                title: "Publicação criada com Sucesso!",
                icon: "success"
            }).then(() => {
                window.location.reload();
            });
        }).fail(function(data) {
            console.log(data);
            Swal.fire({
                title: "Falha ao criar publicação!",
                icon: "error"
            });
        });
        
    });

}
function curtirPost(evento) {
    evento.preventDefault();
    // Aqui vai recuperar o id da publicação
    const idPost = $(this).closest('.card').data('publicacao-id');
    $.ajax({
        url: `/publicacoes/${idPost}/curtir`,
        type: 'POST',
        dataType: 'json',
    }).done(function(data) {
        window.location.reload();
    }).fail(function(data) {
        console.log(data);
        Swal.fire({
            title: "Falha ao curtir publicação!",
            icon: "error"
        });
    });
}
function descurtirPost(evento) {
    evento.preventDefault();
    const idPost = $(this).closest('.card').data('publicacao-id');
    $.ajax({
        url: `/publicacoes/${idPost}/descurtir`,
        type: 'POST',
        dataType: 'json',
    }).done(function(data) {
        window.location.reload();
    }).fail(function(data) {
        console.log(data);
        Swal.fire({
            title: "Falha ao curtir publicação!",
            icon: "error"
        });
    });
}