<template>
  <div class="module">
    <v-card
      v-if="Object.keys(instance).length > 1"
      class="mb-4 pa-2"
      elevation="0"
      color="background"
    >
      <v-row>
        <v-col cols="6">
          <v-text-field
            @change="(newVal) => setValue('title', newVal)"
            label="Name"
            :value="instance.title"
          >
          </v-text-field>
        </v-col>
      </v-row>
      <v-row>
        <v-col cols="6">
          <plans-autocomplete
            :value="instance.billing_plan"
            :custom-params="{
              filters: { type: ['openai'], 'meta.isIndividual': [false] },
              anonymously: false,
            }"
            @input="changeBilling"
            return-object
            label="Price model"
            :rules="planRules"
          />
        </v-col>

        <v-col cols="6" v-if="tarrifAddons.length > 0">
          <v-autocomplete
            @change="(newVal) => setValue('addons', newVal)"
            label="Addons"
            :value="instance.addons"
            :items="isAddonsLoading ? [] : getAvailableAddons()"
            :loading="isAddonsLoading"
            item-value="uuid"
            item-text="title"
            multiple
          />
        </v-col>
      </v-row>
    </v-card>
  </div>
</template>

<script setup>
import { onMounted, ref, toRefs, watch } from "vue";
import useInstanceAddons from "@/hooks/useInstanceAddons";
import plansAutocomplete from "@/components/ui/plansAutoComplete.vue";

const props = defineProps([
  "instance",
  "planRules",
  "spUuid",
  "isEdit",
  "accountId",
]);
const { accountId, instance, isEdit, planRules } = toRefs(props);

const emit = defineEmits(["set-instance", "set-value"]);

const { tarrifAddons, setTariffAddons, getAvailableAddons, isAddonsLoading } =
  useInstanceAddons(instance, (key, value) => setValue(key, value));

const bilingPlan = ref(null);

const getDefaultInstance = () => ({
  title: "instance",
  data: {},
  billing_plan: {},
  addons: [],
  config: {
    user: null,
  },
});

onMounted(() => {
  if (!isEdit.value) {
    emit("set-instance", getDefaultInstance());
  } else {
    changeBilling(instance.value.billing_plan);
  }
  setValue("config.user", accountId.value);
});
const changeBilling = (val) => {
  bilingPlan.value = val;
  setValue("billing_plan", bilingPlan.value);
  setTariffAddons();
};
const setValue = (key, value) => {
  emit("set-value", { key, value });
};

watch(accountId, () => {
  setValue("config.user", accountId.value);
});
</script>

<script>
export default {
  name: "instance-openai-create",
};
</script>

<style scoped></style>
