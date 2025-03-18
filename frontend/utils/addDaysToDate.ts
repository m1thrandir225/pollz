export default function (date: Date, days: number): Date {
  const result = new Date(date);
  result.setDate(result.getDate() + days);

  return result;
}
