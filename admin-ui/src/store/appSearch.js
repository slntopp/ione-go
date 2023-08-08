export default {
  namespaced: true,
  state: {
    searchParam: "",
    variants: {},
    customParams: {},
  },
  mutations: {
    setSearchParam(state, newSearchParam) {
      state.searchParam = newSearchParam;
    },
    setVariants(state, val) {
      state.variants = { ...state.variants, ...val };
    },
    pushVariant(state, { key, value }) {
      state.variants[key] = value;
    },
    resetSearchParams(state) {
      state.searchParam = "";
      state.variants = {};
      state.customParams = {};
    },
    setCustomParam(state, { key, value }) {
      if((key==='searchParam' || !key) && !value.value){
        return
      }
      state.customParams = {
        ...state.customParams,
        [key]: !value.isArray
          ? value
          : [...(state.customParams[key] || []), value],
      };
    },
    deleteCustomParam(state, { key, value, isArray }) {
      if (value && isArray) {
        state.customParams[key] = state.customParams[key].filter(
          (v) => v.value !== value
        );
      } else {
        delete state.customParams[key];
        state.customParams = { ...state.customParams };
      }
    },
  },
  getters: {
    param(state) {
      return state.searchParam;
    },
    variants(state) {
      const variants = { ...state.variants };
      if (Object.keys(variants)) {
        variants["searchParam"] = { title: "Anywhere", key: "searchParam" };
      }
      return variants;
    },
    customParams(state) {
      return state.customParams;
    },
  },
};
