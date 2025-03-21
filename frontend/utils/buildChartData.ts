import type { ChartData } from "chart.js";
import type { PollOptionWithVotes } from "~/types/poll-option";

export default function (
  input: PollOptionWithVotes[],
): ChartData<"doughnut", number[], unknown> {
  const labels = input.map((el) => el.option_text);

  const data = input.map((el) => el.vote_count);

  const colors = data.map(() => randomColor());

  return {
    labels,
    datasets: [
      {
        backgroundColor: colors,
        data,
      },
    ],
  };
}
