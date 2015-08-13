$(document).ready(function() {
  // transform all clickable rows into links
  $('table tr.clickable.row-link').click(function(event) {
    window.location = $(this).data('href');
  });
});