$("#update-publication").on("click", updatePublication);

function updatePublication() {
    $(this).prop("disabled", true);

    const publicationId = $(this).data("publication-id");

    const data = getFormFieldsValues();

    $.ajax({
        url: `/api/publication/${publicationId}`,
        method: "PUT",
        data,
    })
        .done(function () {
            Swal.fire("Sucesso!", "Publição atualizada com sucesso!", "success").then(
                () => {
                    window.location = "/profile";
                }
            );
        })
        .fail(function () {
            Swal.fire("Ops..", "Falha ao atualizar a publicação!", "error");
        })
        .always(() => {
            $("#update-publication").prop("disabled", false);
        });
}

function getFormFieldsValues() {
    const title = $("#title").val();
    const content = $("#content").val();

    return { title, content };
}
