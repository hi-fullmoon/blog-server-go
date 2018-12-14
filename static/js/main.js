$(function () {
    // loading start
    NProgress.start();

    // set scrollbar
    $('body').niceScroll({ cursorwidth: 8, cursorcolor: '#778886' });

    // search
    $('#search-input').on('click', function () {
        $('.js-search-panel').show();
    });

    $('.js-close').on('click', function () {
        $('.js-search-panel').hide();
        $('.js-search-input').val('');
    });

    // set search result container scrollbar
    $('.js-search-result').niceScroll({
        cursorwidth: 8,
        background: '#eee',
        cursorcolor: '#78887d',
        railpadding: { top: 0, right: 0, left: 0, bottom: 0 },
        autohidemode: false,
        cursorborderradius: 0,
        cursorborder: 0
    });

    // loading done
    setTimeout(function () {
        NProgress.done();
    }, 3000);
});