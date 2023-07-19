<template>
    
    <div class="drawer lg:drawer-open">

      <input id="my-drawer-2" type="checkbox" class="drawer-toggle" />

      <div class="drawer-content flex flex-col items-center justify-center">
        <!-- Page content here -->
        
        <label for="my-drawer-2" class="btn btn-primary drawer-button lg:hidden">Open drawer</label>
  
      </div> 

      <div class="drawer-side">
        
            <div class="join">
            <button 
                class="btn join-item"
                @click="showAddModal = true"
              >
                Add
              </button>
                  <button class="btn join-item">TODO</button>
                </div>
        
        <label for="my-drawer-2" class="drawer-overlay"></label>

        <ul class="menu p-4 w-80 h-full bg-base-200 text-base-content">
          <!-- Sidebar content here -->
          <li v-for="(item, index) in items" :key="index">
<!--            <a :href="item.url" target="_blank">{{ item.title }}</a>-->
            <a
                :class="{ 'active': activeFeedId === item.url }"
                @click="activeFeedId = item.url"
            >
              {{ item.title }}</a>
          </li>
        </ul>

      </div>
    </div>

      <!-- 添加组件 -->
      <AddModal 
        v-if="showAddModal"
        @close="showAddModal = false"
        @submit="handleSubmit" 
      />

</template>

<script>
import AddModal from './AddModal.vue'
export default {
    components: {
        AddModal
    },
    data() {
        return {
            showAddModal: false,
          activeFeedId: null,
            items: [
                { title: '信息学院通知', url: 'http://it.bjfu.edu.cn/xyxw/index.html' },
                { title: '材料学院通知', url: 'http://clxy.bjfu.edu.cn/lmdh/xydt/index.html' },
                { title: '教务处通知', url: 'http://jwc.bjfu.edu.cn/ksxx/index.html' },
            ]
        }
    },

    methods: {
        handleSubmit(name, url) {
            // 添加订阅  

            this.showAddModal = false
        }
    }
}
</script>