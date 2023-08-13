<!-- 订阅添加弹窗 -->
<template>    
  <!-- Open the modal using ID.showModal() method -->
  <dialog id="AddSubModal" class="modal">

    <form method="dialog" class="modal-box">
      <h3 class="font-bold text-lg">添加订阅</h3>
      <div class="form-control">
        <label class="label">
          <span class="label-text">显示名称</span>
        </label>
        <input type="text" placeholder="Name" class="input input-bordered" v-model="name" />
      </div>
      <div class="form-control">
        <label class="label">
          <span class="label-text">订阅地址</span>
        </label>
        <input type="text" placeholder="URL" class="input input-bordered" v-model="url" />
      </div>
      <div class="modal-action">
        <button class="btn" @click="submit">提交</button>
        <button class="btn btn-error" @click="$emit('close')">取消</button>
      </div>
    </form>
  </dialog>

</template>


<script>
export default {
    emits: ['close', 'submit'],

    methods: {
        submit() {
            this.$emit('submit', this.name, this.url)
            this.$emit('close')
        }
    }
}  
</script>