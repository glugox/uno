<script setup lang="ts">
import { computed } from 'vue'

export interface Props {
    title: string
    path?: string
}
const props = defineProps<Props>();
const pathEnsured = computed(() => (!props.path ? props.title : props.path))

</script>
<template>
    <div className="page-header d-print-none">
        <div className="row g-2 align-items-center">
            <div className="col">
                <div className="page-pretitle">{{ pathEnsured }}</div>
                <h2 className="page-title">{{ title }}</h2>
            </div>
        </div>
    </div>
</template>