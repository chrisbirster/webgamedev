
widow.addEventListener('load', function() {

	window.addEventListener('load', function() {
		console.log("Hi, Mom!");
		const canvas = document.getElementById('canvas');
		const ctx = canvas.getContext('2d');
		canvas.width = 1280;
		canvas.height = 720; ctx.fillStyle = 'white';
		ctx.font = '48px monospace';
		ctx.textAlign = 'center';
		ctx.textBaseline = 'middle';
		ctx.fillText('Hello, World!', canvas.width / 2, canvas.height / 2);

		// bounce circle off canvas edges taking radius into account
		let x = 50;
		let y = 50;
		let dx = 4;
		let dy = 4;
		let radius = 50;
		let color = 'red';
		function draw() {
			ctx.clearRect(0, 0, canvas.width, canvas.height);
			ctx.beginPath();
			ctx.arc(x, y, radius, 0, Math.PI * 2, true);
			ctx.closePath();
			ctx.fillStyle = color;
			ctx.fill();
			if (x + dx > canvas.width - radius || x + dx < radius) {
				dx = -dx;
				color = 'blue';
			}
			if (y + dy > canvas.height - radius || y + dy < radius) {
				dy = -dy;
				color = 'green';
			}
			x += dx;
			y += dy;
			requestAnimationFrame(draw);
		}
		draw()


	});
