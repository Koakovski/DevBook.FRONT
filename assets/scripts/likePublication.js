const unlikePublicationClass = "unlike-publication";
const likePublicationClass = "like-publication";

$(document).on("click", `.${likePublicationClass}`, likePublication);
$(document).on("click", `.${unlikePublicationClass}`, unlikePublication);

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

            element.addClass(unlikePublicationClass);
            element.addClass("text-danger");
            element.removeClass(likePublicationClass);
        })
        .fail(function () {
            Swal.fire("Ops..", "Falha ao curtir publicação!", "error");
        })
        .always(function () {
            element.prop("disabled", false);
        });
}

function unlikePublication(event) {
    event.preventDefault();

    const element = $(event.target);
    const publicationId = element.closest("div").data("publication-id");

    element.prop("disabled", true);

    $.ajax({
        url: `/api/publication/${publicationId}/unlike`,
        method: "POST",
    })
        .done(function () {
            const likesCounter = element.next("span");
            const likesQuantity = parseInt(likesCounter.text());

            likesCounter.text(likesQuantity - 1);

            element.removeClass(unlikePublicationClass);
            element.removeClass("text-danger");
            element.addClass(likePublicationClass);
        })
        .fail(function () {
            Swal.fire("Ops..", "Falha ao descurtir publicação!", "error");
        })
        .always(function () {
            element.prop("disabled", false);
        });
}
