import type { PollOption } from "./poll-option";

export type Poll = {
  id: string;
  description: string;
  active_until: string;
  created_by: string;
  created_at: string;
  updated_at: string;
};

export type PollWithOptions = {
  poll: Poll;
  options: PollOption[];
};
