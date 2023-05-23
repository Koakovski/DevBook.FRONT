$(".delete-publication").on("click", deletePublication);

function deletePublication(event) {
    console.log(event);
    event.preventDefault();

    const element = $(event.target);
    console.log(element);
    const publication = element.closest("div");
    const publicationId = publication.data("publication-id");

    $.ajax({
        url: `/api/publication/${publicationId}`,
        method: "DELETE",
    })
        .done(function () {
            publication.fadeOut("slow", function () {
                $(this).remove();
            });
        })
        .fail(function (jqXHR) {
            if (jqXHR.status === 200) {
                publication.fadeOut("slow", function () {
                    $(this).remove();
                });
                return;
            }
            alert("Falha!");
        })
        .always(() => {
            $(".delete-publication").prop("disabled", false);
        });
}
