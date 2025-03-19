export type PollOption = {
  id: string;
  poll_id: string;
  option_text: string;
  created_at: string;
};

export type PollOptionWithVotes = {
  id: string;
  poll_id: string;
  option_text: string;
  created_at: string;
  vote_count: number;
};
