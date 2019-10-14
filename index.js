const holidays = require('./holidays.json');

function isHoliday(date) {
	let y = date.getFullYear(), m = date.getMonth()+1, d = date.getDate();
	d = ('0' + d).slice(-2);
	m = ('0' + m).slice(-2);

	return holidays.includes(`${y}/${m}/${d}`);
}

module.exports = isHoliday
