$(".like-publication").on("click", likePublication);

function likePublication(event) {
    event.preventDefault();

    const element = $(event.target);
    const publicationId = element.closest("div").data("publication-id");

    element.prop("disabled", true);

    $.ajax({
        url: `api/publication/${publicationId}/like`,
        method: "POST",
    })
        .done(function () {
            const likesCounter = element.next("span");
            const likesQuantity = parseInt(likesCounter.text());

            likesCounter.text(likesQuantity + 1);
        })
        .fail(function (jqXHR) {
            alert("Falha!");
        })
        .always(function () {
            element.prop("disabled", false);
        });
}
