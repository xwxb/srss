<!-- 订阅添加弹窗 -->
<template>
    <div class="modal-box">
        <form @submit.prevent="submit">
            <!-- 表单 -->
        </form>
    </div>
</template>


<script>
export default {
    emits: ['close', 'submit'],

    methods: {
        submit() {
            this.$emit('submit', name, url)
            this.$emit('close')
        }
    }
}  
</script>