class HandlerAnalytics {
  /*
    @code-graph-label/analytics-pull
    @code-graph-link/analytics-pull,internal-gateway
  */
  executeA() {
    console.log("Execute HandlerAnalyticsPull");
  }

  /*
    @code-graph-label/analytics-push
    @code-graph-link/internal-gateway,analytics-push
  */
  execute() {
    console.log("Execute HandlerAnalyticsPush");
  }
}
