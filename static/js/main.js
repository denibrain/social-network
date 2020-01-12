$(".js-sign-out").click(function () {
    $.post("/signout", {
    }, function (response) {
        document.location = "/";
    }, "json")
        .fail(function(xhr) {
            error("Internal system error")
        });

    return false;
});


$(".user-list > div").click(function () {
    document.location = '/user/' + $(this).data('id')
});