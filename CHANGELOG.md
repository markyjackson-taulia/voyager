# Change Log

## [3.1.0](https://github.com/appscode/voyager/tree/3.1.0) (2017-07-21)
[Full Changelog](https://github.com/appscode/voyager/compare/3.0.0...3.1.0)

**Implemented enhancements:**

- Record events against TPR [\#79](https://github.com/appscode/voyager/issues/79)
- Remove event framework from certificate [\#284](https://github.com/appscode/voyager/pull/284) ([sadlil](https://github.com/sadlil))
- Fix RBAC configs [\#295](https://github.com/appscode/voyager/pull/295) ([sadlil](https://github.com/sadlil))
- Add configure option for Haproxy default timeout. [\#286](https://github.com/appscode/voyager/pull/286) ([sadlil](https://github.com/sadlil))

**Fixed bugs:**

- podAffinityTerm.topologyKey: Required value: can not be empty [\#320](https://github.com/appscode/voyager/issues/320)
- Restore objects if deleted by mistake. [\#283](https://github.com/appscode/voyager/issues/283)
- HostPort mode does not work for AWS [\#281](https://github.com/appscode/voyager/issues/281)
- Externalservice redirection gets reset [\#279](https://github.com/appscode/voyager/issues/279)
- Voyager doesn't work with cloud = minikube and type = HostPort [\#272](https://github.com/appscode/voyager/issues/272)
- Adding cert manager to existing ingress does not open port 443 [\#267](https://github.com/appscode/voyager/issues/267)
- Bug: annotations are not applied [\#266](https://github.com/appscode/voyager/issues/266)
- Add newline in pem file [\#261](https://github.com/appscode/voyager/issues/261)
- Adding SSL to an existing ingress does not mount certs [\#260](https://github.com/appscode/voyager/issues/260)
- Set topology key for pod anti-affinity [\#321](https://github.com/appscode/voyager/pull/321) ([tamalsaha](https://github.com/tamalsaha))
- Correctly detect changed ports [\#322](https://github.com/appscode/voyager/pull/322) ([tamalsaha](https://github.com/tamalsaha))
- Fix Adding SSL to an existing ingress does not mount certs \#260 [\#306](https://github.com/appscode/voyager/pull/306) ([sadlil](https://github.com/sadlil))
- Fix External Service redirect Issue [\#304](https://github.com/appscode/voyager/pull/304) ([sadlil](https://github.com/sadlil))
- Fix RBAC configs [\#295](https://github.com/appscode/voyager/pull/295) ([sadlil](https://github.com/sadlil))
- Fix Operator panic on service restore [\#273](https://github.com/appscode/voyager/pull/273) ([sadlil](https://github.com/sadlil))

**Closed issues:**

- Difficulties to setup, scarce docs [\#303](https://github.com/appscode/voyager/issues/303)
- Setup Issues [\#298](https://github.com/appscode/voyager/issues/298)
- Setup Issues [\#297](https://github.com/appscode/voyager/issues/297)
- Install Voyager as critical addon [\#292](https://github.com/appscode/voyager/issues/292)
- configurable HAProxy defaults [\#280](https://github.com/appscode/voyager/issues/280)
- Support setting resource for pods [\#277](https://github.com/appscode/voyager/issues/277)
- The link to contribution guide in README.md is broken. [\#274](https://github.com/appscode/voyager/issues/274)
- Voyager exporter sidecar isn't exporting any metrics [\#270](https://github.com/appscode/voyager/issues/270)
- Adding an AWS Cert and opening 80 and 443 doesn't work for plain http:// [\#268](https://github.com/appscode/voyager/issues/268)
- Support HorizontalPodAutoscaling for HAProxy pods [\#242](https://github.com/appscode/voyager/issues/242)
- Test updated chart with RBAC [\#302](https://github.com/appscode/voyager/issues/302)
- Delete TPR when NS is deleted [\#258](https://github.com/appscode/voyager/issues/258)
- voyager-operator should ensure that ServiceAccount/Role/RoleBinding exists for created voyager deploys. [\#252](https://github.com/appscode/voyager/issues/252)
- RBAC objects for Voyager operator. [\#241](https://github.com/appscode/voyager/issues/241)
- Should all hosts be passed to EnsureLoadBalancer [\#88](https://github.com/appscode/voyager/issues/88)

**Merged pull requests:**

- Fix various chart issues [\#324](https://github.com/appscode/voyager/pull/324) ([tamalsaha](https://github.com/tamalsaha))
- Add Custom timeout docs [\#323](https://github.com/appscode/voyager/pull/323) ([sadlil](https://github.com/sadlil))
- Revendor dependencies. [\#312](https://github.com/appscode/voyager/pull/312) ([tamalsaha](https://github.com/tamalsaha))
- move RecognizeWellKnownRegions\(\) to the beginning of newAWSCloud\(\) [\#311](https://github.com/appscode/voyager/pull/311) ([jipperinbham](https://github.com/jipperinbham))
- Add ingress label to exported metrics [\#300](https://github.com/appscode/voyager/pull/300) ([tamalsaha](https://github.com/tamalsaha))
- Support setting resource for pods [\#289](https://github.com/appscode/voyager/pull/289) ([tamalsaha](https://github.com/tamalsaha))
- fix the contribution guild link \(\#274\) [\#275](https://github.com/appscode/voyager/pull/275) ([aimof](https://github.com/aimof))
- Update aws-cert-manager.md [\#269](https://github.com/appscode/voyager/pull/269) ([julianvmodesto](https://github.com/julianvmodesto))
- Add command reference docs [\#265](https://github.com/appscode/voyager/pull/265) ([tamalsaha](https://github.com/tamalsaha))
- Point to HPA example on readme pages. [\#254](https://github.com/appscode/voyager/pull/254) ([tamalsaha](https://github.com/tamalsaha))
- Add example with hpa [\#253](https://github.com/appscode/voyager/pull/253) ([julianvmodesto](https://github.com/julianvmodesto))
- Use ```console instead of ```sh syntax highlighting [\#309](https://github.com/appscode/voyager/pull/309) ([tamalsaha](https://github.com/tamalsaha))
- Install Voyager as critical addon [\#301](https://github.com/appscode/voyager/pull/301) ([tamalsaha](https://github.com/tamalsaha))
- Add Stats Service events [\#299](https://github.com/appscode/voyager/pull/299) ([sadlil](https://github.com/sadlil))
- Recover ServiceMonitor [\#294](https://github.com/appscode/voyager/pull/294) ([tamalsaha](https://github.com/tamalsaha))
- Make node selectors optional for HostPort [\#293](https://github.com/appscode/voyager/pull/293) ([tamalsaha](https://github.com/tamalsaha))
- Delete kube lister classes. [\#291](https://github.com/appscode/voyager/pull/291) ([tamalsaha](https://github.com/tamalsaha))
- Record events against TPR [\#290](https://github.com/appscode/voyager/pull/290) ([tamalsaha](https://github.com/tamalsaha))
- Add tpr constants [\#288](https://github.com/appscode/voyager/pull/288) ([tamalsaha](https://github.com/tamalsaha))
- Remove event framework [\#282](https://github.com/appscode/voyager/pull/282) ([tamalsaha](https://github.com/tamalsaha))
- Update dev docs. [\#264](https://github.com/appscode/voyager/pull/264) ([tamalsaha](https://github.com/tamalsaha))
- Add a newline between crt & key. [\#263](https://github.com/appscode/voyager/pull/263) ([tamalsaha](https://github.com/tamalsaha))
- Create RBAC roles for Voyager during installation [\#256](https://github.com/appscode/voyager/pull/256) ([tamalsaha](https://github.com/tamalsaha))
- Support non-default service account with offshoot pods [\#255](https://github.com/appscode/voyager/pull/255) ([tamalsaha](https://github.com/tamalsaha))

## [3.0.0](https://github.com/appscode/voyager/tree/3.0.0) (2017-06-23)
[Full Changelog](https://github.com/appscode/voyager/compare/1.5.6...3.0.0)

**Implemented enhancements:**

- Automatically create ServiceMonitor for built-in exporter [\#154](https://github.com/appscode/voyager/issues/154)
- Fix testframework for aws and update docs. [\#237](https://github.com/appscode/voyager/pull/237) ([sadlil](https://github.com/sadlil))

**Fixed bugs:**

- Delete pods & services matching old labels before starting operator [\#229](https://github.com/appscode/voyager/issues/229)
- Check for updates properly [\#250](https://github.com/appscode/voyager/pull/250) ([tamalsaha](https://github.com/tamalsaha))
- Don't restore stats service if stats is disabled. [\#249](https://github.com/appscode/voyager/pull/249) ([tamalsaha](https://github.com/tamalsaha))
- Apply labels to stats service for service monitor [\#248](https://github.com/appscode/voyager/pull/248) ([tamalsaha](https://github.com/tamalsaha))
- Fix Bugs [\#247](https://github.com/appscode/voyager/pull/247) ([sadlil](https://github.com/sadlil))
- Correctly parse target port [\#245](https://github.com/appscode/voyager/pull/245) ([tamalsaha](https://github.com/tamalsaha))
- Fix testframework for aws and update docs. [\#237](https://github.com/appscode/voyager/pull/237) ([sadlil](https://github.com/sadlil))
- Add dns-resovler-check-health annotation to for ExternalName service [\#226](https://github.com/appscode/voyager/pull/226) ([tamalsaha](https://github.com/tamalsaha))
- Add cloud config file  [\#218](https://github.com/appscode/voyager/pull/218) ([sadlil](https://github.com/sadlil))
- Fix bugs [\#217](https://github.com/appscode/voyager/pull/217) ([sadlil](https://github.com/sadlil))

**Closed issues:**

- Add chart value for --cloud-config mount [\#228](https://github.com/appscode/voyager/issues/228)
- Document http-\>https redirect with AWS cert manager [\#225](https://github.com/appscode/voyager/issues/225)
- Update version policy [\#194](https://github.com/appscode/voyager/issues/194)
- Change api group to voyager.appscode.com [\#193](https://github.com/appscode/voyager/issues/193)
- Use client-go [\#192](https://github.com/appscode/voyager/issues/192)
- Use pod anti-affinity for deployments [\#161](https://github.com/appscode/voyager/issues/161)
- Change api group to voyager.appscode.com [\#142](https://github.com/appscode/voyager/issues/142)

**Merged pull requests:**

- Small typo fix \(CLOUDE\_CONFIG =\> CLOUD\_CONFIG\) [\#251](https://github.com/appscode/voyager/pull/251) ([thecodeassassin](https://github.com/thecodeassassin))
- Document http-\>https redirect with AWS cert manager [\#235](https://github.com/appscode/voyager/pull/235) ([tamalsaha](https://github.com/tamalsaha))
- Remove deprecated Daemon type. [\#205](https://github.com/appscode/voyager/pull/205) ([tamalsaha](https://github.com/tamalsaha))
- Automatically create ServiceMonitor for built-in exporter [\#203](https://github.com/appscode/voyager/pull/203) ([tamalsaha](https://github.com/tamalsaha))
- Track operator version [\#200](https://github.com/appscode/voyager/pull/200) ([tamalsaha](https://github.com/tamalsaha))
- Update version policy to point to client-go [\#198](https://github.com/appscode/voyager/pull/198) ([tamalsaha](https://github.com/tamalsaha))
- Use client-go [\#196](https://github.com/appscode/voyager/pull/196) ([tamalsaha](https://github.com/tamalsaha))
- Use stats service port name in ServiceMonitor [\#246](https://github.com/appscode/voyager/pull/246) ([tamalsaha](https://github.com/tamalsaha))
- Use correct api schema when checking ingress class. [\#244](https://github.com/appscode/voyager/pull/244) ([tamalsaha](https://github.com/tamalsaha))
- Note test-ns policy [\#243](https://github.com/appscode/voyager/pull/243) ([tamalsaha](https://github.com/tamalsaha))
- Add acs  provider [\#236](https://github.com/appscode/voyager/pull/236) ([tamalsaha](https://github.com/tamalsaha))
- Update chart readme for cloud config [\#234](https://github.com/appscode/voyager/pull/234) ([tamalsaha](https://github.com/tamalsaha))
- Make cloud config configurable. [\#233](https://github.com/appscode/voyager/pull/233) ([tamalsaha](https://github.com/tamalsaha))
- Change api group to networking.appscode.com [\#232](https://github.com/appscode/voyager/pull/232) ([tamalsaha](https://github.com/tamalsaha))
- Update \*\*\*Getter interfaces match form [\#231](https://github.com/appscode/voyager/pull/231) ([tamalsaha](https://github.com/tamalsaha))
- Delete pods & services matching old labels before starting operator [\#230](https://github.com/appscode/voyager/pull/230) ([tamalsaha](https://github.com/tamalsaha))
- Use PreRun & PostRun to send analytics. [\#224](https://github.com/appscode/voyager/pull/224) ([tamalsaha](https://github.com/tamalsaha))
- Update metric endpoints documentation. [\#223](https://github.com/appscode/voyager/pull/223) ([tamalsaha](https://github.com/tamalsaha))
- Fix port used for exposing metrics from operator. [\#222](https://github.com/appscode/voyager/pull/222) ([tamalsaha](https://github.com/tamalsaha))
- Open both port 443 & 80 when AWS cert manager is in use. [\#221](https://github.com/appscode/voyager/pull/221) ([tamalsaha](https://github.com/tamalsaha))
- Mount cloud config in chart [\#220](https://github.com/appscode/voyager/pull/220) ([tamalsaha](https://github.com/tamalsaha))
- Use root user inside docker [\#219](https://github.com/appscode/voyager/pull/219) ([tamalsaha](https://github.com/tamalsaha))
- Rename exporter port to targetPort [\#216](https://github.com/appscode/voyager/pull/216) ([tamalsaha](https://github.com/tamalsaha))
- Use Voyager group name correctly. [\#215](https://github.com/appscode/voyager/pull/215) ([tamalsaha](https://github.com/tamalsaha))
- Update default ports [\#214](https://github.com/appscode/voyager/pull/214) ([tamalsaha](https://github.com/tamalsaha))
- Update docs for service monitor integration [\#213](https://github.com/appscode/voyager/pull/213) ([tamalsaha](https://github.com/tamalsaha))
- Fix unit test build issues [\#210](https://github.com/appscode/voyager/pull/210) ([tamalsaha](https://github.com/tamalsaha))
- Change api group to voyager.appscode.com [\#209](https://github.com/appscode/voyager/pull/209) ([tamalsaha](https://github.com/tamalsaha))
- Update docs to point to 3.0.0 [\#208](https://github.com/appscode/voyager/pull/208) ([tamalsaha](https://github.com/tamalsaha))
- Stop creating stats service. [\#207](https://github.com/appscode/voyager/pull/207) ([tamalsaha](https://github.com/tamalsaha))
- Update labels applied to HAProxy pods & services. [\#206](https://github.com/appscode/voyager/pull/206) ([tamalsaha](https://github.com/tamalsaha))
- Fix client-go fake import [\#204](https://github.com/appscode/voyager/pull/204) ([tamalsaha](https://github.com/tamalsaha))
- Change default HAProxy image to 1.7.6-3.0.0 [\#202](https://github.com/appscode/voyager/pull/202) ([tamalsaha](https://github.com/tamalsaha))
- Add HAProxy 1.7.6 dockerfiles. [\#201](https://github.com/appscode/voyager/pull/201) ([tamalsaha](https://github.com/tamalsaha))
- Add voyager export command [\#199](https://github.com/appscode/voyager/pull/199) ([tamalsaha](https://github.com/tamalsaha))
- Only keep Firewall\(\) interface in cloud provider [\#195](https://github.com/appscode/voyager/pull/195) ([tamalsaha](https://github.com/tamalsaha))

## [1.5.6](https://github.com/appscode/voyager/tree/1.5.6) (2017-06-16)
[Full Changelog](https://github.com/appscode/voyager/compare/1.5.5...1.5.6)

**Implemented enhancements:**

- Delete docker image from docker hub after integration test [\#125](https://github.com/appscode/voyager/issues/125)
- Change how stats work [\#106](https://github.com/appscode/voyager/issues/106)
- Use AWS ELB Proxy Protocol [\#100](https://github.com/appscode/voyager/issues/100)
- Track Kube's refactoring cloud provider API [\#36](https://github.com/appscode/voyager/issues/36)
- Expose HAProxy stats to prometheus [\#13](https://github.com/appscode/voyager/issues/13)
- Support AWS cert manager [\#189](https://github.com/appscode/voyager/pull/189) ([tamalsaha](https://github.com/tamalsaha))
- Merge existing pods and service during create ingress resource [\#181](https://github.com/appscode/voyager/pull/181) ([sadlil](https://github.com/sadlil))
- Add support for ServiceTypeExternalName [\#167](https://github.com/appscode/voyager/pull/167) ([sadlil](https://github.com/sadlil))
- Collect analytics for voyager usages [\#133](https://github.com/appscode/voyager/pull/133) ([sadlil](https://github.com/sadlil))
- Fix stats behavior [\#130](https://github.com/appscode/voyager/pull/130) ([sadlil](https://github.com/sadlil))
- Improve test framework [\#121](https://github.com/appscode/voyager/pull/121) ([sadlil](https://github.com/sadlil))

**Fixed bugs:**

- Error out if Daemon type does not provide a node selector. [\#159](https://github.com/appscode/voyager/issues/159)
- Disable analytics when running tests [\#147](https://github.com/appscode/voyager/issues/147)
- Missing services should be an warning not error stack error [\#137](https://github.com/appscode/voyager/issues/137)
- Bad ingress object results in unstable HAProxy [\#135](https://github.com/appscode/voyager/issues/135)
- Add ingress hostname to Ingress [\#132](https://github.com/appscode/voyager/issues/132)
- Deleting LB deployment does not get recreated [\#123](https://github.com/appscode/voyager/issues/123)
- Ensure HAproxy running when endpoints changes. [\#120](https://github.com/appscode/voyager/issues/120)
- Updating Ingress annotations are not picked up by controller [\#115](https://github.com/appscode/voyager/issues/115)
- Fix Ingress Status Update Properly. [\#134](https://github.com/appscode/voyager/pull/134) ([sadlil](https://github.com/sadlil))
- Expose monitoring port in chart and deploy yamls [\#156](https://github.com/appscode/voyager/pull/156) ([tamalsaha](https://github.com/tamalsaha))
- Add LoadBalancerSourceRange to ingress Spec [\#148](https://github.com/appscode/voyager/pull/148) ([sadlil](https://github.com/sadlil))
- Ensure loadbalancer resource [\#145](https://github.com/appscode/voyager/pull/145) ([sadlil](https://github.com/sadlil))
- Add annotation to add accept-proxy in bind statements [\#144](https://github.com/appscode/voyager/pull/144) ([sadlil](https://github.com/sadlil))
- Remove unwanted stacktrace from log [\#139](https://github.com/appscode/voyager/pull/139) ([sadlil](https://github.com/sadlil))
- Fix stats behavior [\#130](https://github.com/appscode/voyager/pull/130) ([sadlil](https://github.com/sadlil))

**Closed issues:**

- Allow exposing port 443 on the LoadBalancer Service [\#188](https://github.com/appscode/voyager/issues/188)
- Source IP detection  [\#146](https://github.com/appscode/voyager/issues/146)
- helm chart [\#113](https://github.com/appscode/voyager/issues/113)
- Merge pods & services even on create [\#172](https://github.com/appscode/voyager/issues/172)
- Document 1.5.6 changes [\#150](https://github.com/appscode/voyager/issues/150)
- Support Services of type ExternalName [\#127](https://github.com/appscode/voyager/issues/127)
- Collect usage analytics [\#126](https://github.com/appscode/voyager/issues/126)
- Support use of field spec.loadBalancerSourceRanges on Services of type LoadBalancer [\#122](https://github.com/appscode/voyager/issues/122)

**Merged pull requests:**

- Fix chart path [\#191](https://github.com/appscode/voyager/pull/191) ([tamalsaha](https://github.com/tamalsaha))
-  ./hack/make.py test\_deploy to generate deploymnts yaml [\#184](https://github.com/appscode/voyager/pull/184) ([ashiquzzaman33](https://github.com/ashiquzzaman33))
- Disable analytics for test runs [\#182](https://github.com/appscode/voyager/pull/182) ([tamalsaha](https://github.com/tamalsaha))
- Prepare docs for 1.5.6 [\#178](https://github.com/appscode/voyager/pull/178) ([tamalsaha](https://github.com/tamalsaha))
- Remove cluster name flag [\#177](https://github.com/appscode/voyager/pull/177) ([tamalsaha](https://github.com/tamalsaha))
- Remove persist annotation [\#174](https://github.com/appscode/voyager/pull/174) ([tamalsaha](https://github.com/tamalsaha))
- Add TLS certs for testing [\#173](https://github.com/appscode/voyager/pull/173) ([tamalsaha](https://github.com/tamalsaha))
- Run kloader check without exec [\#171](https://github.com/appscode/voyager/pull/171) ([tamalsaha](https://github.com/tamalsaha))
- Error out if Daemon type does not provide a node selector. [\#168](https://github.com/appscode/voyager/pull/168) ([tamalsaha](https://github.com/tamalsaha))
- Remove dependency on k8s-addons [\#141](https://github.com/appscode/voyager/pull/141) ([tamalsaha](https://github.com/tamalsaha))
- Use kloader 1.5.1 and check config before starting runit. [\#140](https://github.com/appscode/voyager/pull/140) ([tamalsaha](https://github.com/tamalsaha))
- Use ci-space cluster for testing [\#131](https://github.com/appscode/voyager/pull/131) ([ashiquzzaman33](https://github.com/ashiquzzaman33))
- tcp.md: fix typo/port mismatch [\#119](https://github.com/appscode/voyager/pull/119) ([alekssaul](https://github.com/alekssaul))
- Add Jenkinsfile [\#118](https://github.com/appscode/voyager/pull/118) ([ashiquzzaman33](https://github.com/ashiquzzaman33))
- Jenkins test patch1 [\#117](https://github.com/appscode/voyager/pull/117) ([ashiquzzaman33](https://github.com/ashiquzzaman33))
- Document flag options [\#190](https://github.com/appscode/voyager/pull/190) ([tamalsaha](https://github.com/tamalsaha))
- Docs for 1.5.6 [\#183](https://github.com/appscode/voyager/pull/183) ([sadlil](https://github.com/sadlil))
- Set metrics port to :8080 by default [\#180](https://github.com/appscode/voyager/pull/180) ([tamalsaha](https://github.com/tamalsaha))
- Stop redefining -h flag for run command. [\#179](https://github.com/appscode/voyager/pull/179) ([tamalsaha](https://github.com/tamalsaha))
- Remove --cluster-name flag [\#176](https://github.com/appscode/voyager/pull/176) ([tamalsaha](https://github.com/tamalsaha))
- Add nil check before reading options from Ingress annotations. [\#170](https://github.com/appscode/voyager/pull/170) ([tamalsaha](https://github.com/tamalsaha))
- Various cleanup of annotations [\#169](https://github.com/appscode/voyager/pull/169) ([tamalsaha](https://github.com/tamalsaha))
- Use hyphen separated words as annotation key. [\#166](https://github.com/appscode/voyager/pull/166) ([tamalsaha](https://github.com/tamalsaha))
- Use ingress.appscode.com/keep-source-ip: true to preserve source IP [\#165](https://github.com/appscode/voyager/pull/165) ([tamalsaha](https://github.com/tamalsaha))
- Combine annotation keys ip & persist into persist [\#162](https://github.com/appscode/voyager/pull/162) ([tamalsaha](https://github.com/tamalsaha))
- Make nodeSelector annotation applicable for any mode. [\#160](https://github.com/appscode/voyager/pull/160) ([tamalsaha](https://github.com/tamalsaha))
- Explain versioning policy. [\#158](https://github.com/appscode/voyager/pull/158) ([tamalsaha](https://github.com/tamalsaha))
- Apply various comments from official charts team [\#157](https://github.com/appscode/voyager/pull/157) ([tamalsaha](https://github.com/tamalsaha))
- Move component docs directly under user-guide [\#155](https://github.com/appscode/voyager/pull/155) ([tamalsaha](https://github.com/tamalsaha))
- Expose Operator & HAProxy metrics [\#153](https://github.com/appscode/voyager/pull/153) ([tamalsaha](https://github.com/tamalsaha))
- Reorganize code to add run sub command [\#152](https://github.com/appscode/voyager/pull/152) ([tamalsaha](https://github.com/tamalsaha))
- Add forked cloudprovider in third\_party package [\#151](https://github.com/appscode/voyager/pull/151) ([tamalsaha](https://github.com/tamalsaha))

## [1.5.5](https://github.com/appscode/voyager/tree/1.5.5) (2017-05-22)
[Full Changelog](https://github.com/appscode/voyager/compare/1.5.4...1.5.5)

**Implemented enhancements:**

- Support user provided annotations [\#103](https://github.com/appscode/voyager/issues/103)
- Rename Daemon type to HostPort [\#72](https://github.com/appscode/voyager/issues/72)
- expose NodePort like functionality to Ingress [\#68](https://github.com/appscode/voyager/issues/68)
- Cross Namespace Service Support [\#40](https://github.com/appscode/voyager/issues/40)
- Support health checks [\#38](https://github.com/appscode/voyager/issues/38)
- Support full spectrum of HAProxy rules [\#21](https://github.com/appscode/voyager/issues/21)
- Add user provided annotations in LoadBalancer in Service/Pods [\#105](https://github.com/appscode/voyager/pull/105) ([sadlil](https://github.com/sadlil))
- Feature weighted backend [\#77](https://github.com/appscode/voyager/pull/77) ([sadlil](https://github.com/sadlil))
- Update svc instead of Deleting svc [\#87](https://github.com/appscode/voyager/pull/87) ([sadlil](https://github.com/sadlil))
- Feature: backend rules [\#80](https://github.com/appscode/voyager/pull/80) ([sadlil](https://github.com/sadlil))

**Fixed bugs:**

- Update service in NodePort & LoadBalancer mode [\#86](https://github.com/appscode/voyager/issues/86)
- Fix ALPN negotiation [\#32](https://github.com/appscode/voyager/issues/32)
- Use annotations for backend weight [\#83](https://github.com/appscode/voyager/pull/83) ([sadlil](https://github.com/sadlil))
- Fix Loadbalancer Port Open Issues [\#99](https://github.com/appscode/voyager/pull/99) ([sadlil](https://github.com/sadlil))
- Ensure pod delete [\#97](https://github.com/appscode/voyager/pull/97) ([sadlil](https://github.com/sadlil))
- Update svc instead of Deleting svc [\#87](https://github.com/appscode/voyager/pull/87) ([sadlil](https://github.com/sadlil))

**Closed issues:**

- Allow free form HTTP rewriting [\#76](https://github.com/appscode/voyager/issues/76)
- Test NodePort mode [\#98](https://github.com/appscode/voyager/issues/98)
- Ensure pods are deleted before deleting RC / Deployment [\#96](https://github.com/appscode/voyager/issues/96)
- Test that previously open NodePort is not reassigned [\#95](https://github.com/appscode/voyager/issues/95)
- Use HAProxy 1.7.5 [\#90](https://github.com/appscode/voyager/issues/90)
- Document 1.5.5 milestone features [\#78](https://github.com/appscode/voyager/issues/78)
- Specify different services in a backend with weights [\#75](https://github.com/appscode/voyager/issues/75)

**Merged pull requests:**

- Update top readme file [\#112](https://github.com/appscode/voyager/pull/112) ([tamalsaha](https://github.com/tamalsaha))
- Update docs [\#111](https://github.com/appscode/voyager/pull/111) ([tamalsaha](https://github.com/tamalsaha))
- NodePort Tests, Annotations Documentation [\#110](https://github.com/appscode/voyager/pull/110) ([sadlil](https://github.com/sadlil))
- Change HAProxy image to 1.7.5-1.5.5 [\#93](https://github.com/appscode/voyager/pull/93) ([tamalsaha](https://github.com/tamalsaha))
- Rename Daemon type to HostPort [\#84](https://github.com/appscode/voyager/pull/84) ([tamalsaha](https://github.com/tamalsaha))
- Use appscode/errors v2 [\#81](https://github.com/appscode/voyager/pull/81) ([tamalsaha](https://github.com/tamalsaha))
- Avoid upgrade in operator docker image [\#109](https://github.com/appscode/voyager/pull/109) ([tamalsaha](https://github.com/tamalsaha))
- Use alpine as the base image for operator [\#107](https://github.com/appscode/voyager/pull/107) ([tamalsaha](https://github.com/tamalsaha))
- Add `go` and `glide` commands to developer docs [\#101](https://github.com/appscode/voyager/pull/101) ([julianvmodesto](https://github.com/julianvmodesto))
- Ensure forward secrecy [\#94](https://github.com/appscode/voyager/pull/94) ([tamalsaha](https://github.com/tamalsaha))
- Update docs to build HAProxy 1.7.5 [\#92](https://github.com/appscode/voyager/pull/92) ([tamalsaha](https://github.com/tamalsaha))
- Use HAProxy 1.7.5 [\#91](https://github.com/appscode/voyager/pull/91) ([tamalsaha](https://github.com/tamalsaha))
- Introduce NodePort mode [\#85](https://github.com/appscode/voyager/pull/85) ([tamalsaha](https://github.com/tamalsaha))
- Update 1.5.5 Documentation [\#82](https://github.com/appscode/voyager/pull/82) ([sadlil](https://github.com/sadlil))

## [1.5.4](https://github.com/appscode/voyager/tree/1.5.4) (2017-05-08)
[Full Changelog](https://github.com/appscode/voyager/compare/1.5.3...1.5.4)

**Fixed bugs:**

- Voyager pod is restarting itself when attached backend pod restarts [\#69](https://github.com/appscode/voyager/issues/69)
- Do not restart lb pod when backend pod restarts [\#70](https://github.com/appscode/voyager/pull/70) ([sadlil](https://github.com/sadlil))

**Merged pull requests:**

- Rename operator deployment to voyager-operator [\#71](https://github.com/appscode/voyager/pull/71) ([tamalsaha](https://github.com/tamalsaha))

## [1.5.3](https://github.com/appscode/voyager/tree/1.5.3) (2017-05-03)
[Full Changelog](https://github.com/appscode/voyager/compare/1.5.2...1.5.3)

**Implemented enhancements:**

- Support StatefulSet pod names in Voyager [\#14](https://github.com/appscode/voyager/issues/14)
- Ingress Hostname based traffic forwarding [\#66](https://github.com/appscode/voyager/pull/66) ([sadlil](https://github.com/sadlil))

**Fixed bugs:**

- cloud-provider & cloud-name can't be always required [\#64](https://github.com/appscode/voyager/issues/64)

**Merged pull requests:**

- Prepare docs for 1.5.3 release [\#67](https://github.com/appscode/voyager/pull/67) ([tamalsaha](https://github.com/tamalsaha))
- cloud-provider & cloud-name is not required for unknown providers. [\#65](https://github.com/appscode/voyager/pull/65) ([tamalsaha](https://github.com/tamalsaha))
- Test/fix ingress name [\#63](https://github.com/appscode/voyager/pull/63) ([ashiquzzaman33](https://github.com/ashiquzzaman33))
- Update docs to new chart location [\#60](https://github.com/appscode/voyager/pull/60) ([tamalsaha](https://github.com/tamalsaha))
- Move chart to root directory [\#59](https://github.com/appscode/voyager/pull/59) ([tamalsaha](https://github.com/tamalsaha))

## [1.5.2](https://github.com/appscode/voyager/tree/1.5.2) (2017-04-21)
[Full Changelog](https://github.com/appscode/voyager/compare/1.5.1...1.5.2)

**Implemented enhancements:**

- Add Retry on DaemonMode Loadbalancer http test call [\#52](https://github.com/appscode/voyager/pull/52) ([sadlil](https://github.com/sadlil))
- Fix Documentation [\#51](https://github.com/appscode/voyager/pull/51) ([sadlil](https://github.com/sadlil))

**Fixed bugs:**

- Slack channel token\_revoked [\#48](https://github.com/appscode/voyager/issues/48)
- Service ports should be int [\#47](https://github.com/appscode/voyager/issues/47)

**Merged pull requests:**

- Add service to deployments.yaml [\#58](https://github.com/appscode/voyager/pull/58) ([tamalsaha](https://github.com/tamalsaha))
- Prepare docs for version 1.5.2 [\#57](https://github.com/appscode/voyager/pull/57) ([tamalsaha](https://github.com/tamalsaha))
- Add service in voyager [\#56](https://github.com/appscode/voyager/pull/56) ([saumanbiswas](https://github.com/saumanbiswas))
- Fix stable chart [\#55](https://github.com/appscode/voyager/pull/55) ([saumanbiswas](https://github.com/saumanbiswas))
- Use unversioned time. [\#54](https://github.com/appscode/voyager/pull/54) ([tamalsaha](https://github.com/tamalsaha))
- Doc/fix update [\#53](https://github.com/appscode/voyager/pull/53) ([sadlil](https://github.com/sadlil))
- Initial voyager chart [\#43](https://github.com/appscode/voyager/pull/43) ([saumanbiswas](https://github.com/saumanbiswas))

## [1.5.1](https://github.com/appscode/voyager/tree/1.5.1) (2017-04-05)
[Full Changelog](https://github.com/appscode/voyager/compare/1.5.0...1.5.1)

**Implemented enhancements:**

- Enable GKE [\#44](https://github.com/appscode/voyager/issues/44)

**Merged pull requests:**

- Enable GKE [\#45](https://github.com/appscode/voyager/pull/45) ([tamalsaha](https://github.com/tamalsaha))
- Fix Typos [\#42](https://github.com/appscode/voyager/pull/42) ([sunkuet02](https://github.com/sunkuet02))
- update README [\#41](https://github.com/appscode/voyager/pull/41) ([ben-st](https://github.com/ben-st))

## [1.5.0](https://github.com/appscode/voyager/tree/1.5.0) (2017-03-01)
**Implemented enhancements:**

- Various clean ups [\#18](https://github.com/appscode/voyager/issues/18)
- Add ALPN options to TCP Backends [\#35](https://github.com/appscode/voyager/pull/35) ([sadlil](https://github.com/sadlil))
- Update docs with voyager options and test modes [\#34](https://github.com/appscode/voyager/pull/34) ([sadlil](https://github.com/sadlil))
- Add alpn option while TLS is used [\#25](https://github.com/appscode/voyager/pull/25) ([sadlil](https://github.com/sadlil))
- Adding Tests - Unit and E2E [\#12](https://github.com/appscode/voyager/pull/12) ([sadlil](https://github.com/sadlil))
- Ensure TPR at runtime [\#9](https://github.com/appscode/voyager/pull/9) ([sadlil](https://github.com/sadlil))
- add ingress-class [\#4](https://github.com/appscode/voyager/pull/4) ([sadlil](https://github.com/sadlil))
- Renamed ingress annotations to "ingress.appscode.com" [\#3](https://github.com/appscode/voyager/pull/3) ([sadlil](https://github.com/sadlil))
- use updated reloader. [\#2](https://github.com/appscode/voyager/pull/2) ([sadlil](https://github.com/sadlil))

**Fixed bugs:**

- Failing to deploy [\#29](https://github.com/appscode/voyager/issues/29)
- Remove ALPN h2 for https [\#33](https://github.com/appscode/voyager/pull/33) ([sadlil](https://github.com/sadlil))
- Update doc fix for \#19 [\#26](https://github.com/appscode/voyager/pull/26) ([sadlil](https://github.com/sadlil))

**Closed issues:**

- Update documentation for nodeSelector cleanup [\#24](https://github.com/appscode/voyager/issues/24)

**Merged pull requests:**

- Add doc explaining release process. [\#37](https://github.com/appscode/voyager/pull/37) ([tamalsaha](https://github.com/tamalsaha))
- Pass KLOADER\_ARGS as env variable [\#30](https://github.com/appscode/voyager/pull/30) ([tamalsaha](https://github.com/tamalsaha))
- Init cloud provider for Azure. [\#28](https://github.com/appscode/voyager/pull/28) ([tamalsaha](https://github.com/tamalsaha))
- Revendor dependencies. [\#23](https://github.com/appscode/voyager/pull/23) ([tamalsaha](https://github.com/tamalsaha))
- Use Ubuntu:16.04 as the base image to enable ALPN. [\#22](https://github.com/appscode/voyager/pull/22) ([tamalsaha](https://github.com/tamalsaha))
- Resolve \#18 [\#19](https://github.com/appscode/voyager/pull/19) ([sadlil](https://github.com/sadlil))
- Add example on front page. [\#16](https://github.com/appscode/voyager/pull/16) ([tamalsaha](https://github.com/tamalsaha))
- README typos [\#15](https://github.com/appscode/voyager/pull/15) ([JakeAustwick](https://github.com/JakeAustwick))
- Add links of subsections [\#11](https://github.com/appscode/voyager/pull/11) ([sadlil](https://github.com/sadlil))
- Update docs [\#10](https://github.com/appscode/voyager/pull/10) ([sadlil](https://github.com/sadlil))
- Rename voyager to Voyager  [\#8](https://github.com/appscode/voyager/pull/8) ([sadlil](https://github.com/sadlil))
- Add acknowledgements [\#7](https://github.com/appscode/voyager/pull/7) ([sadlil](https://github.com/sadlil))
- Documentation for voyager [\#6](https://github.com/appscode/voyager/pull/6) ([sadlil](https://github.com/sadlil))
- Use kloader. [\#5](https://github.com/appscode/voyager/pull/5) ([tamalsaha](https://github.com/tamalsaha))
- Custom pongo2 filters for loading haproxy data. [\#1](https://github.com/appscode/voyager/pull/1) ([sadlil](https://github.com/sadlil))



\* *This Change Log was automatically generated by [github_changelog_generator](https://github.com/skywinder/Github-Changelog-Generator)*