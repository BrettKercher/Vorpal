$( document ).ready(function() {

	$('#player-one-div').hide();
	$('#player-two-div').hide();
	$('#player-three-div').hide();
	$('#player-four-div').hide();

	$('.player-count button').click(function () {
		$('#player-count-holder').val($(this).text());
	});

	// Player Count Clicked

	$('#players-two').click(function() {
		$('#player-count-holder').val('2');
		$('#player-one-div').show();
		$('#player-two-div').show();
		$('#player-three-div').hide();
		$('#player-four-div').hide();

	});

	$('#players-three').click(function() {
		$('#player-count-holder').val('3');
		$('#player-one-div').show();
		$('#player-two-div').show();
		$('#player-three-div').show();
		$('#player-four-div').hide();
	});

	$('#players-four').click(function() {
		$('#player-count-holder').val('4');
		$('#player-one-div').show();
		$('#player-two-div').show();
		$('#player-three-div').show();
		$('#player-four-div').show();
	});


	// Is Random Clicked

	$('#isP1Random').change(function() {
		if($('#isP1Random').is(':checked')) {
			$('#playerOneRandom').val("1");
		} else {
			$('#playerOneRandom').val("0");
		}
	});

	$('#isP2Random').change(function() {
		if($('#isP2Random').is(':checked')) {
			$('#playerTwoRandom').val("1");
		} else {
			$('#playerTwoRandom').val("0");
		}
	});

	$('#isP3Random').change(function() {
		if($('#isP3Random').is(':checked')) {
			$('#playerThreeRandom').val("1");
		} else {
			$('#playerThreeRandom').val("0");
		}
	});

	$('#isP4Random').change(function() {
		if($('#isP4Random').is(':checked')) {
			$('#playerFourRandom').val("1");
		} else {
			$('#playerFourRandom').val("0");
		}
	});

	//Player 1 Stats

	$('#p1KillsMinus').click(function() {
		var value = parseInt($('#p1Kills').val());
		var newVal = (value - 1);
		$('#p1Kills').val(newVal);
	});

	$('#p1KillsPlus').click(function() {
		var value = parseInt($('#p1Kills').val());
		var newVal = (value + 1);
		$('#p1Kills').val(newVal);
	});

	$('#p1DeathsMinus').click(function() {
		var value = parseInt($('#p1Deaths').val());
		var newVal = (value - 1);
		$('#p1Deaths').val(newVal);
	});

	$('#p1DeathsPlus').click(function() {
		var value = parseInt($('#p1Deaths').val());
		var newVal = (value + 1);
		$('#p1Deaths').val(newVal);
	});

	$('#p1SDsMinus').click(function() {
		var value = parseInt($('#p1SDs').val());
		var newVal = (value - 1);
		$('#p1SDs').val(newVal);
	});

	$('#p1SDsPlus').click(function() {
		var value = parseInt($('#p1SDs').val());
		var newVal = (value + 1);
		$('#p1SDs').val(newVal);
	});

	//Player 2 Stats

	$('#p2KillsMinus').click(function() {
		var value = parseInt($('#p2Kills').val());
		var newVal = (value - 1);
		$('#p2Kills').val(newVal);
	});

	$('#p2KillsPlus').click(function() {
		var value = parseInt($('#p2Kills').val());
		var newVal = (value + 1);
		$('#p2Kills').val(newVal);
	});

	$('#p2DeathsMinus').click(function() {
		var value = parseInt($('#p2Deaths').val());
		var newVal = (value - 1);
		$('#p2Deaths').val(newVal);
	});

	$('#p2DeathsPlus').click(function() {
		var value = parseInt($('#p2Deaths').val());
		var newVal = (value + 1);
		$('#p2Deaths').val(newVal);
	});

	$('#p2SDsMinus').click(function() {
		var value = parseInt($('#p2SDs').val());
		var newVal = (value - 1);
		$('#p2SDs').val(newVal);
	});

	$('#p2SDsPlus').click(function() {
		var value = parseInt($('#p2SDs').val());
		var newVal = (value + 1);
		$('#p2SDs').val(newVal);
	});

	//Player 3 Stats

	$('#p3KillsMinus').click(function() {
		var value = parseInt($('#p3Kills').val());
		var newVal = (value - 1);
		$('#p3Kills').val(newVal);
	});

	$('#p3KillsPlus').click(function() {
		var value = parseInt($('#p3Kills').val());
		var newVal = (value + 1);
		$('#p3Kills').val(newVal);
	});

	$('#p3DeathsMinus').click(function() {
		var value = parseInt($('#p3Deaths').val());
		var newVal = (value - 1);
		$('#p3Deaths').val(newVal);
	});

	$('#p3DeathsPlus').click(function() {
		var value = parseInt($('#p3Deaths').val());
		var newVal = (value + 1);
		$('#p3Deaths').val(newVal);
	});

	$('#p3SDsMinus').click(function() {
		var value = parseInt($('#p3SDs').val());
		var newVal = (value - 1);
		$('#p3SDs').val(newVal);
	});

	$('#p3SDsPlus').click(function() {
		var value = parseInt($('#p3SDs').val());
		var newVal = (value + 1);
		$('#p3SDs').val(newVal);
	});

	//Player 4 Stats

	$('#p4KillsMinus').click(function() {
		var value = parseInt($('#p4Kills').val());
		var newVal = (value - 1);
		$('#p4Kills').val(newVal);
	});

	$('#p4KillsPlus').click(function() {
		var value = parseInt($('#p4Kills').val());
		var newVal = (value + 1);
		$('#p4Kills').val(newVal);
	});

	$('#p4DeathsMinus').click(function() {
		var value = parseInt($('#p4Deaths').val());
		var newVal = (value - 1);
		$('#p4Deaths').val(newVal);
	});

	$('#p4DeathsPlus').click(function() {
		var value = parseInt($('#p4Deaths').val());
		var newVal = (value + 1);
		$('#p4Deaths').val(newVal);
	});

	$('#p4SDsMinus').click(function() {
		var value = parseInt($('#p4SDs').val());
		var newVal = (value - 1);
		$('#p4SDs').val(newVal);
	});

	$('#p4SDsPlus').click(function() {
		var value = parseInt($('#p4SDs').val());
		var newVal = (value + 1);
		$('#p4SDs').val(newVal);
	});

});