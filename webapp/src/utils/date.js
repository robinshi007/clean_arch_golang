export default function formatDate(date) {
  const diff = new Date() - date; // the difference in milliseconds

  if (diff < 1000) { // less than 1 second
    return 'right now';
  }

  const sec = Math.floor(diff / 1000); // convert diff to seconds

  if (sec < 60) {
    return `${sec} sec. ago`;
  }

  const min = Math.floor(diff / 60000); // convert diff to minutes
  if (min < 60) {
    return `${min} min. ago`;
  }

  // format the date
  // add leading zeroes to single-digit day/month/hours/minutes
  let d = date;
  d = [
    `${d.getFullYear()}`,
    `0${d.getMonth() + 1}`,
    `0${d.getDate()}`,
    `0${d.getHours()}`,
    `0${d.getMinutes()}`,
    `0${d.getSeconds()}`,
  ].map(component => component.slice(-2)); // take last 2 digits of every component

  // join the components into date
  return `${d.slice(0, 3).join('/')} ${d.slice(3).join(':')}`;
}
