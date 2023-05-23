$(".delete-publication").on("click", deletePublication);

function deletePublication(event) {
    event.preventDefault();

    Swal.fire({
        title: "Atenção!",
        text: "Tem certeza que deseja excluir essa publicação, essa ação é irreversível!",
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "warning",
    }).then(function (confirmation) {
        if (!confirmation.isConfirmed) return;

        handleDelete(event);
    });
}

function handleDelete(event) {
    const element = $(event.target);
    const publication = element.closest("div");
    const publicationId = publication.data("publication-id");

    $.ajax({
        url: `/api/publication/${publicationId}`,
        method: "DELETE",
    })
        .done(function () {
            handleDeleteSucess(publication);
        })
        .fail(function (jqXHR) {
            if (jqXHR.status === 200) {
                handleDeleteSucess(publication);
                return;
            }
            Swal.fire("Ops...", "Falha ao excluir a publicação!", "error");
        })
        .always(() => {
            $(".delete-publication").prop("disabled", false);
        });
}

function handleDeleteSucess(publication) {
    Swal.fire("Sucesso!", "Publição excluida com sucesso!", "success").then(() => {
        publication.fadeOut("slow", function () {
            $(this).remove();
        });
    });
}
