<template>

  <!-- 抽屉 -->
  <div class="drawer lg:drawer-open">

    // 抽屉开关
    <input id="my-drawer-2" type="checkbox" class="drawer-toggle" />

    // 主要内容
    <div class="drawer-content flex flex-col items-center justify-center">
      <!-- Page content here -->
      <div class="card w-2/3 bg-green-100 p-4 shadow" v-for="news in news[0]">
        <div class="card-body">
          <h2 class="card-title font-bold text-lg mb-2">{{ news.title }}</h2>
          <p>{{ news.content }}</p>
        </div>
      </div>

      <label for="my-drawer-2" class="btn btn-primary drawer-button lg:hidden">Open drawer</label>

    </div>

    // 侧栏内容
    <div class="drawer-side">

      <div class="join">
        <button class="btn join-item" onclick="AddSubModal.showModal()">Add</button>
        <button class="btn join-item">TODO</button>
      </div>

      <label for="my-drawer-2" class="drawer-overlay"></label>

      <ul class="menu p-4 w-80 h-full bg-base-200 text-base-content">
        <!-- Sidebar content here -->
        <li v-for="(item, index) in items" :key="index">
          <!--            <a :href="item.url" target="_blank">{{ item.title }}</a>-->
          <a :class="{ 'active': activeFeedId === item.url }" @click="activeFeedId = item.url">
            {{ item.title }}</a>
        </li>
      </ul>

    </div>

  </div>

  <!-- 添加组件 -->
  <AddSubModal @form-submit="handleSubmit" />
</template>

<script>
import AddSubModal from './AddSubModal.vue'
export default {
  components: {
    AddSubModal
  },
  data() {
    return {
      activeFeedId: null,
      items: [
        { title: '信息学院通知', url: 'http://it.bjfu.edu.cn/xyxw/index.html' },
        { title: '材料学院通知', url: 'http://clxy.bjfu.edu.cn/lmdh/xydt/index.html' },
        { title: '教务处通知', url: 'http://jwc.bjfu.edu.cn/ksxx/index.html' },
      ],
      news: [
        [
          { title: '关于提交2022-2023学年第二学期课程成绩的通知', content: '      各学院、体育教学部：2022-2023学年第二学期已接近尾声，为便于学生及家长尽早查询到课程成绩，保障后续各项教学工作的有序进行，现将本学期成绩提交工作安排通知如下：1.本学期成绩提交截止时间为2023年7月30日。2.各学院、体育教学部应加强对试卷评阅和成绩提交的管理。要及时关注本单位教师的成绩提交情况，提醒任课教师尽早、按时提交课程成绩，方便学生查询成绩。3.建议优先使用360浏览器兼容模式登陆教务系统录入成绩，成绩提交后应及时检查录入结果。4.教师在录入成绩过程中若遇到问题，请及时联系学院教学秘书或教务处。教务处联系人：曹晓庆    联系电话:62338099  特此通知。             北京林业大学教务处2023年7月3日  版权所有 © 2004-2012    北京林业大学教务处    E-mail: jwc@bjfu.edu.cn ' },
          { title: '22-23-2学期《汽车技术新发展》考试通知', content: '考试班级：车辆22级-1.2.3班级考试时间：2023年6月15日（周四）上午9：00～11：00考试地点：二教 201版权所有 © 2004-2012    北京林业大学教务处    E-mail: jwc@bjfu.edu.cn ' }
        ]
      ]
    }
  },

  methods: {
    handleSubmit(formData) {
      // 添加订阅  
      this.items.push({ title: formData.name, url: formData.url })
      
    },
    grabDataFromUrl(url) {
      // 从url中抓取数据

    }
  }
}
</script>