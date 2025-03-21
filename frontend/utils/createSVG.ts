import { renderSVG } from "uqr";

export default function (input: string): string {
  return renderSVG(input, {
    border: 2,
    ecc: "L",
  });
}
