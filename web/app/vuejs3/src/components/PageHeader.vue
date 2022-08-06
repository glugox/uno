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
    <div class="page-header mb-3">
        <div class="row align-items-center mw-100">
            <div class="col">
                <div class="mb-1">
                    <ol class="breadcrumb" aria-label="breadcrumbs">
                        <li class="breadcrumb-item">
                            <a href="#">Administration</a>
                        </li>

                        <li class="breadcrumb-item active" aria-current="page">
                            <a href="#">{{ pathEnsured }}</a>
                        </li>
                    </ol>
                </div>
                <h2 class="page-title">
                    <span class="text-truncate">{{ title }}</span>
                </h2>
            </div>
            <div class="col-auto">
                <div class="btn-list">
                    <a href="#" class="btn btn-primary"> Save Settings </a>
                </div>
            </div>
        </div>
    </div>

</template>



