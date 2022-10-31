import api from "@/api.js";

export default {
  namespaced: true,
  state: {
    namespaces: [],
    loading: false,
  },
  mutations: {
    setNamespaces(state, namespaces) {
      state.namespaces = namespaces;
    },
    setLoading(state, data) {
      state.loading = data;
    },
  },
  actions: {
    fetch({ commit }) {
      commit("setLoading", true);
      return new Promise((resolve, reject) => {
        api.namespaces
          .list()
          .then((response) => {
            commit("setNamespaces", response.pool);
            resolve(response);
          })
          .catch((error) => {
            reject(error);
          })
          .finally(() => {
            commit("setLoading", false);
          });
      });
    },
  },
  getters: {
    all(state) {
      return state.namespaces;
    },
    isLoading(state) {
      return state.loading;
    },
  },
};
