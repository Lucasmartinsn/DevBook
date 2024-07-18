$('#seguir-usuario').on('click', seguirUser);
$('#deixar-seguir-usuario').on('click', deixarSeguirUser);

function seguirUser(evento) {
    evento.preventDefault();
    var id = $(this).closest('.card').data('usuario-id');
    console.log("seguindo..... Id do usuario seguido: " + id);
    $.ajax({
        url: `/usuario/${id}/seguir`,
        type: 'GET',
        dataType: 'json',
    }).done(function (data) {
        Swal.fire({
            title: "Agora voce esta seguindo esse usuario!",
            icon: "success"
        }).then(() => {
            window.location.reload();
        });
    }).fail(function (data) {
        console.log(data.responseJSON);
        Swal.fire({
            title: "Falha ao seguir esse Usuario!",
            icon: "error"
        });
    });

}

function deixarSeguirUser(evento) {
    evento.preventDefault();
    var id = $(this).closest('.card').data('usuario-id');

    Swal.fire({
        title: "Atenção!",
        text: "Você quer deixar de seguir esse Usuario?",
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "warning"
    }).then((confirmacao) => {
        if (!confirmacao.isConfirmed) return;

        $.ajax({
            url: `/usuario/${id}/deixar-de-seguir`,
            type: 'DELETE',
            dataType: 'json',
        }).done(function (data) {
            Swal.fire({
                title: "Você deixou de seguir esse usuario!",
                icon: "success"
            }).then(() => {
                window.location.reload();
            });
        }).fail(function (data) {
            console.log(data.responseJSON);
            Swal.fire({
                title: "Falha ao deixar de seguir!",
                icon: "error"
            });
        });

    });
}