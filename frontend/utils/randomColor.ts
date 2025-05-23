export default function (): string {
  const colors: string[] = [
    "#f87171",
    "#fb923c",
    "#fbbf24",
    "#facc15",
    "#a3e635",
    "#4ade80",
    "#34d399",
    "#2dd4bf",
    "#22d3ee",
    "#38bdf8",
    "#60a5fa",
    "#818cf8",
    "#a78bfa",
    "#c084fc",
    "#e879f9",
    "#f472b6",
    "#fb7185",
    "#f472b6",
  ];

  const random = Math.floor(Math.random() * colors.length);
  return colors[random];
}
