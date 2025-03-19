import type { Vote } from "~/types/vote";

/**
 * Client-side only holds all the votes has had
 */
export const useVoteStore = defineStore(
  "votes",
  () => {
    const votes = ref<Vote[]>([]);

    function findVoteById(voteId: string) {
      return votes.value.filter((el) => el.id === voteId);
    }

    function findVoteByOptionId(optionId: string) {
      return votes.value.filter((el) => el.option_id === optionId);
    }

    function addVote(vote: Vote) {
      votes.value.push(vote);
    }

    function clearVotes() {
      votes.value = [];
    }

    function removeVote(voteId: string) {
      votes.value.map((el) => el.id != voteId);
    }

    return {
      votes,
      addVote,
      removeVote,
      clearVotes,
      findVoteById,
      findVoteByOptionId,
    };
  },
  {
    persist: {
      storage: piniaPluginPersistedstate.cookies(),
    },
  },
);
