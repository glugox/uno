<script lang="ts" setup>
import { mapActions, storeToRefs } from "pinia";
import { onMounted, ref } from "vue";
import { useCounterStore } from "@/store/modules/example";
import { usePokemonStore } from "@/store/modules/pokemon";

import PageHeader from "@/components/PageHeader.vue";

const welcome = ref("OK");

// Instance to store
const main = useCounterStore();
const pokemon = usePokemonStore();
// Make data reactive
const { counter, doubleCounter } = storeToRefs(main);
const { pokemonsName } = storeToRefs(pokemon);
// Mapping actions
const { increment } = mapActions(useCounterStore, ["increment"]);
// Reset store data
const reset = () => main.$reset();
// Call action from store to get pokemons on mounted lifecycle
onMounted(() => pokemon.getPokemons());
</script>

<template>
  <PageHeader title="Settings" />
  <h1>{{ welcome }}</h1>
  <h3>Counter using Pinia Store</h3>
  <p>Counter: {{ counter }}</p>
  <p>Double counter: {{ doubleCounter }}</p>
  <button @click="increment">Increment</button>
  <button @click="reset">Reset store</button>
  <div v-for="(poke, i) of pokemonsName" :key="i">{{ poke }}</div>
</template>
