export default function (date: Date, weeks: number): Date {
  const result = new Date(date);
  result.setDate(result.getDate() + weeks * 7);

  return result;
}
