<template>
  <div class="problem">
    <el-row>
      <el-col :span="22" :offset="1">
        <el-tabs
          v-model="activeTab"
          type="card"
          @tab-remove="removeTab"
          :before-leave="beforeLeave"
        >
          <el-tab-pane
            closable
            v-for="item in tabs"
            :key="item.name"
            :label="item.title"
            :name="item.name"
            :lazy="false"
            ><problemTab
              v-on:title="problemTitle"
              v-bind:pid="pid"
              v-bind:oj="oj"
              :username="name"
              ref="child"
            ></problemTab>
          </el-tab-pane>
          <el-tab-pane key="add" name="add">
            <div
              slot="label"
              style="font-size: 15px"
              class="el-icon-plus"
            ></div>
          </el-tab-pane>
        </el-tabs>
      </el-col>
    </el-row>
  </div>
</template>
 
<script>
import problemTab from "@/components/problemTab";
export default {
  name: "problem",
  components: {
    problemTab,
  },
  data() {
    return {
      activeTab: "0",
      tabs: [],
      tabIndex: 0,
      name: "LLLLLL0420",
      pid: "1000",
      oj: "HDU",
    };
  },
  mounted: function () {
    this.addTab();
  },
  methods: {
    toProblem: function (oj, pid) {
      this.addTab();
      this.oj = oj;
      this.pid = pid;
    },

    problemTitle: function (childValue) {
      this.tabs.forEach((item) => {
        if (item.name == this.activeTab) {
          item.title = childValue;
        }
      });
    },
    beforeLeave(currentName, oldName) {
      //重点，如果name是add，则什么都不触发
      if (currentName == "add") {
        this.addTab();
        return false;
      } else {
        this.currentIndex = currentName;
        return true;
      }
    },
    addTab() {
      let newTabName = ++this.tabIndex + "";
      this.tabs.push({
        title: "问题" + this.tabIndex,
        name: newTabName,
        content: problemTab,
      });
      this.activeTab = newTabName;
    },
    removeTab(targetName) {
      let tabs = this.tabs;
      let activeName = this.activeTab;
      if (activeName === targetName) {
        tabs.forEach((tab, index) => {
          if (tab.name === targetName) {
            let nextTab = tabs[index + 1] || tabs[index - 1];
            if (nextTab) {
              activeName = nextTab.name;
            }
          }
        });
      }
      this.activeTab = activeName;
      this.tabs = tabs.filter((tab) => tab.name !== targetName);
    },
  },
};
</script>


<style>
</style>