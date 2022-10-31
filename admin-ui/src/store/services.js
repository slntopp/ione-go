import api from "@/api.js";

export default {
  namespaced: true,
  state: {
    services: [],
    service: [],
    loading: false,
    loadingItem: false,
  },
  getters: {
    all(state) {
      return state.services;
    },
    one(state) {
      return state.service;
    },
    isLoading(state) {
      return state.loading;
    },
    isLoadingItem(state) {
      return state.loadingItem;
    },
  },
  mutations: {
    setServices(state, services) {
      state.services = services;
    },
    setLoading(state, data) {
      state.loading = data;
    },
    setLoadingItem(state, data) {
      state.loadingItem = data;
    },
    setService(state, service) {
      if (state.service.length) {
        let isProductExists = false;
        state.service.find((item) => {
          if (item.uuid === service.uuid) {
            isProductExists = true;
          }
        });
        if (!isProductExists) {
          state.service.push(service);
        }
      } else {
        state.service.push(service);
      }
    },
    updateService(state, service) {
      if (!state.services.length) state.services.push(service);
      state.services = state.services.map((serv) =>
        serv.uuid === service.uuid ? service : serv
      );
    },
    updateInstance(state, { value, uuid }) {
      const i = state.services.findIndex((el) => uuid === el.uuid);
      const service = state.services[i];

      service.instancesGroups.forEach((el, i, groups) => {
        el.instances.forEach(({ uuid }, j) => {
          if (uuid === value.uuid) {
            groups[i].instances[j].state = value.state;
          }
        });
      });
    },
    fetchByIdElem(state, data) {
      state.service = data;
    },
  },
  actions: {
    fetch({ commit }) {
      commit("setLoading", true);
      return new Promise((resolve, reject) => {
        api.services
          .list()
          .then((response) => {
            const servicesWithoutDel = response.pool.filter(
              (s) => s.status !== "DEL"
            );
            commit("setServices", servicesWithoutDel);
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
    fetchById({ commit }, id) {
      commit("setLoading", true);
      return new Promise((resolve, reject) => {
        api.services
          .get(id)
          .then((response) => {
            commit("updateService", response);
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
    fetchByIdItem({ commit }, id) {
      commit("setLoadingItem", true);
      return new Promise((resolve, reject) => {
        api.services
          .get(id)
          .then((response) => {
            commit("setService", response);
            resolve(response);
          })
          .catch((error) => {
            reject(error);
          })
          .finally(() => {
            commit("setLoadingItem", false);
          });
      });
    },
  },
};
