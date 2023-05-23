$("#follow").on("click", followUser);
$("#unfollow").on("click", unfollowUser);

function followUser() {
    const userId = $(this).data("user-id");

    $(this).prop("disabled", true);

    $.ajax({
        url: `/api/user/${userId}/follow`,
        method: "POST",
    })
        .done(function () {
            window.location = `/user/${userId}`;
        })
        .fail(function () {
            Swal.fire("Ops...", "Falha ao seguir usuário a publicação!", "error");
        })
        .always(() => {
            $("#follow").prop("disabled", false);
        });
}

function unfollowUser() {
    const userId = $(this).data("user-id");

    $(this).prop("disabled", true);

    $.ajax({
        url: `/api/user/${userId}/unfollow`,
        method: "POST",
    })
        .done(function () {
            window.location = `/user/${userId}`;
        })
        .fail(function () {
            Swal.fire(
                "Ops...",
                "Falha ao parar de seguir usuário a publicação!",
                "error"
            );
        })
        .always(() => {
            $("#unfollow").prop("disabled", false);
        });
}
