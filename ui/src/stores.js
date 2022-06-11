import { readable, writable } from "svelte/store";

const createWritableStore = (key, startValue) => {
  const { subscribe, set } = writable(startValue);

  return {
    subscribe,
    set,
    useLocalStorage: () => {
      const json = localStorage.getItem(key);
      if (json) {
        set(JSON.parse(json));
      }

      subscribe((current) => {
        localStorage.setItem(key, JSON.stringify(current));
      });
    },
  };
};

export const User = (function () {
  const { subscribe, set, useLocalStorage } = createWritableStore(
    "user",
    false
  );
  return {
    subscribe,
    signout: () => {
      set(false);
    },
    signin: (agent_name) => {
      set(agent_name);
    },
    useLocalStorage,
  };
})();

export const ActiveAccount = (function () {
  const { subscribe, set, useLocalStorage } = createWritableStore(
    "active_account",
    false
  );
  return {
    subscribe,
    set,
    useLocalStorage,
  };
})();

export const BaseURL = readable("http://localhost:80/api/v1");
