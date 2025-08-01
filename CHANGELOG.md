# Changelog

## 0.8.0 (2025-08-01)

Full Changelog: [v0.7.0-alpha...v0.8.0](https://github.com/m3ter-com/m3ter-sdk-go/compare/v0.7.0-alpha...v0.8.0)

### Features

* **api:** remove audit fields from Terraform provider ([d0e1796](https://github.com/m3ter-com/m3ter-sdk-go/commit/d0e17960a8467e4c7cf28f954c824d3c3c3797b9))

## 0.7.0-alpha (2025-07-28)

Full Changelog: [v0.6.2-alpha...v0.7.0-alpha](https://github.com/m3ter-com/m3ter-sdk-go/compare/v0.6.2-alpha...v0.7.0-alpha)

### Features

* **api:** add aggregationId query param to ListPricings ([4c79b28](https://github.com/m3ter-com/m3ter-sdk-go/commit/4c79b2818d576ac4b822ee55bf7984524b877d29))


### Bug Fixes

* **client:** bad casing ([3cbb265](https://github.com/m3ter-com/m3ter-sdk-go/commit/3cbb26546a083d926214de791dac972c85606395))
* **client:** process custom base url ahead of time ([f53eb05](https://github.com/m3ter-com/m3ter-sdk-go/commit/f53eb05e6eda906a6f9ceeaf3675f004f7997e36))

## 0.6.2-alpha (2025-07-18)

Full Changelog: [v0.6.1-alpha...v0.6.2-alpha](https://github.com/m3ter-com/m3ter-sdk-go/compare/v0.6.1-alpha...v0.6.2-alpha)

### Bug Fixes

* **api:** use WebhookDestinationResponse in webhook endpoint response schemas ([8ca98ca](https://github.com/m3ter-com/m3ter-sdk-go/commit/8ca98ca990b6ea2be660aaa740e1029413b1a657))


### Chores

* lint tests in subpackages ([eaccbd0](https://github.com/m3ter-com/m3ter-sdk-go/commit/eaccbd056302afda952c5373c553d58e7419e344))

## 0.6.1-alpha (2025-07-09)

Full Changelog: [v0.6.0-alpha...v0.6.1-alpha](https://github.com/m3ter-com/m3ter-sdk-go/compare/v0.6.0-alpha...v0.6.1-alpha)

## 0.6.0-alpha (2025-07-09)

Full Changelog: [v0.5.0-alpha...v0.6.0-alpha](https://github.com/m3ter-com/m3ter-sdk-go/compare/v0.5.0-alpha...v0.6.0-alpha)

### Features

* **api:** mark `version` attribute as computed in Terraform ([83c91de](https://github.com/m3ter-com/m3ter-sdk-go/commit/83c91de16e6d72389c93128e42131513dd219d1e))


### Bug Fixes

* don't try to deserialize as json when ResponseBodyInto is []byte ([f115ace](https://github.com/m3ter-com/m3ter-sdk-go/commit/f115ace2c4dc4bb42ffa2394c2ce3e479b74a54f))
* **pagination:** check if page data is empty in GetNextPage ([2fccbd5](https://github.com/m3ter-com/m3ter-sdk-go/commit/2fccbd5fa802e011ac1eb0d3f8c477f4bcdbb5e4))


### Chores

* **ci:** only run for pushes and fork pull requests ([c6b514d](https://github.com/m3ter-com/m3ter-sdk-go/commit/c6b514d6c717e359e7903b2fd0da25a7a9d7fabc))
* **internal:** codegen related update ([77adec8](https://github.com/m3ter-com/m3ter-sdk-go/commit/77adec82db36084875e4cedf93af00e929898b3c))
* **internal:** fix lint script for tests ([1c56c85](https://github.com/m3ter-com/m3ter-sdk-go/commit/1c56c852dabb3bf359865c2516b414d960338c24))
* lint tests ([52e470e](https://github.com/m3ter-com/m3ter-sdk-go/commit/52e470e9680261e4bfdc6b0239c13135952ad570))

## 0.5.0-alpha (2025-06-24)

Full Changelog: [v0.4.0-alpha...v0.5.0-alpha](https://github.com/m3ter-com/m3ter-sdk-go/compare/v0.4.0-alpha...v0.5.0-alpha)

### Features

* **api:** Set "version" attribute to "computed_optional" in Terraform ([04417ca](https://github.com/m3ter-com/m3ter-sdk-go/commit/04417ca67566330927a9ec701c3c2e139013a9f6))
* **client:** add debug log helper ([76a6950](https://github.com/m3ter-com/m3ter-sdk-go/commit/76a6950de118edfcd3ec56b6c87b5ba38d3e3d34))


### Bug Fixes

* **api:** terraform release readiness ([6626b97](https://github.com/m3ter-com/m3ter-sdk-go/commit/6626b97f0227a56a54bf798c95dd43bfdadcfb36))


### Chores

* **ci:** enable for pull requests ([974c0c5](https://github.com/m3ter-com/m3ter-sdk-go/commit/974c0c5d78c766362caf9abe25cfdcb3357be3e9))

## 0.4.0-alpha (2025-06-11)

Full Changelog: [v0.3.0-alpha...v0.4.0-alpha](https://github.com/m3ter-com/m3ter-sdk-go/compare/v0.3.0-alpha...v0.4.0-alpha)

### Features

* **api:** add statements ([288a991](https://github.com/m3ter-com/m3ter-sdk-go/commit/288a991a1bf7fa4c29e0645154ba97bba8781b96))
* **api:** Introduce OrganizationConfigRequest model ([86abb22](https://github.com/m3ter-com/m3ter-sdk-go/commit/86abb224b6a161fb54ab8c533c55c2f001d9f9d7))
* **api:** OpenAPI spec update ([c244ebf](https://github.com/m3ter-com/m3ter-sdk-go/commit/c244ebf0b1fbb62c86bea95d1c11a64165dd3098))
* **api:** Spec fixes ([55338d7](https://github.com/m3ter-com/m3ter-sdk-go/commit/55338d75c25039f160be9ef3a228edd6b3903ac5))
* **api:** update open api spec ([9804ae8](https://github.com/m3ter-com/m3ter-sdk-go/commit/9804ae80ae6e9ebeb9770bce4e992c23ff1997f4))
* **api:** update OpenAPI spec + associated fixes ([acc6452](https://github.com/m3ter-com/m3ter-sdk-go/commit/acc64529b30652d0792b2bfcad7ce7d7ef30a760))
* **client:** add support for endpoint-specific base URLs in python ([c1220bc](https://github.com/m3ter-com/m3ter-sdk-go/commit/c1220bccdbc09f22cbcca0942f414c05d5efa4ac))
* **client:** add support for reading base URL from environment variable ([6ba1213](https://github.com/m3ter-com/m3ter-sdk-go/commit/6ba1213335259211eb748076bdccf385763928fd))


### Bug Fixes

* **client:** clean up reader resources ([c911398](https://github.com/m3ter-com/m3ter-sdk-go/commit/c911398ec044e5e21de92d440a2fc4711145d023))
* **client:** correctly update body in WithJSONSet ([c47b975](https://github.com/m3ter-com/m3ter-sdk-go/commit/c47b975e606af4138d9b144f1e94bf94bf18b1b4))
* handle empty bodies in WithJSONSet ([da1a180](https://github.com/m3ter-com/m3ter-sdk-go/commit/da1a180754c30a7b9711d1727af46f90c36fa3e8))
* **pagination:** handle errors when applying options ([c369b29](https://github.com/m3ter-com/m3ter-sdk-go/commit/c369b296f448202bd89b172d8052179d9693de35))


### Chores

* **ci:** add timeout thresholds for CI jobs ([9733540](https://github.com/m3ter-com/m3ter-sdk-go/commit/9733540df33ef5cee5f26d6dbca04f33db064c06))
* **ci:** only use depot for staging repos ([f3c016a](https://github.com/m3ter-com/m3ter-sdk-go/commit/f3c016a18ea859b0769dc84bf50cc3b1cab19905))
* **ci:** run on more branches and use depot runners ([6889134](https://github.com/m3ter-com/m3ter-sdk-go/commit/6889134847367a5febceeed8199881d04ffa9f3b))
* **docs:** document pre-request options ([60fc814](https://github.com/m3ter-com/m3ter-sdk-go/commit/60fc81418b06ab0732397a253fca0114d7e22c00))
* **docs:** grammar improvements ([a0b5238](https://github.com/m3ter-com/m3ter-sdk-go/commit/a0b5238af703b9d8e78b9051d580dfa39f305aa0))
* improve devcontainer setup ([14af707](https://github.com/m3ter-com/m3ter-sdk-go/commit/14af7073d1abfdec1453ab657a8b2b7e60b0fcf2))
* **internal:** codegen related update ([092b3de](https://github.com/m3ter-com/m3ter-sdk-go/commit/092b3deaaae6fc20b7a668be5bf17a6d7532d24e))
* **internal:** updates ([d2f02a1](https://github.com/m3ter-com/m3ter-sdk-go/commit/d2f02a184b199b36c22269d0ab58afb256051d89))
* make go mod tidy continue on error ([c5b0bd8](https://github.com/m3ter-com/m3ter-sdk-go/commit/c5b0bd8763ab3b773af586c15a4778dadda33046))


### Documentation

* update documentation links to be more uniform ([b65641e](https://github.com/m3ter-com/m3ter-sdk-go/commit/b65641eda07a4404729d6c9c2665cb5cbaec16c9))

## 0.3.0-alpha (2025-04-10)

Full Changelog: [v0.2.0-alpha...v0.3.0-alpha](https://github.com/m3ter-com/m3ter-sdk-go/compare/v0.2.0-alpha...v0.3.0-alpha)

### Features

* **api:** add measurement request model ([a163e11](https://github.com/m3ter-com/m3ter-sdk-go/commit/a163e11c5dacaf9784d070b7addbec58a7914653))
* fixup! Use ingest.m3ter.com hostname for usage ingest endpoint ([33f9176](https://github.com/m3ter-com/m3ter-sdk-go/commit/33f91763a8ef929940eed9b2816199d6d52aa919))
* Use ingest.m3ter.com hostname for usage ingest endpoint ([00a9573](https://github.com/m3ter-com/m3ter-sdk-go/commit/00a95732fabcc4cc2f8a1fc90053c8cbadec28fb))

## 0.2.0-alpha (2025-04-10)

Full Changelog: [v0.1.0-alpha...v0.2.0-alpha](https://github.com/m3ter-com/m3ter-sdk-go/compare/v0.1.0-alpha...v0.2.0-alpha)

### Features

* **api:** rename DataFieldResponse to DataField and add DerivedField as a model ([58049ea](https://github.com/m3ter-com/m3ter-sdk-go/commit/58049eaf9100a189c841eb0cdbfbeeca10d5ade3))


### Chores

* **internal:** expand CI branch coverage ([#85](https://github.com/m3ter-com/m3ter-sdk-go/issues/85)) ([c9b3f9f](https://github.com/m3ter-com/m3ter-sdk-go/commit/c9b3f9fc66efc7a93d5ff9bf3d3373e860504ac5))
* **internal:** reduce CI branch coverage ([86fcb85](https://github.com/m3ter-com/m3ter-sdk-go/commit/86fcb85b80b28f28f6df8cd4552f8bbffb136acf))

## 0.1.0-alpha (2025-04-09)

Full Changelog: [v0.1.0-alpha.14...v0.1.0-alpha](https://github.com/m3ter-com/m3ter-sdk-go/compare/v0.1.0-alpha.14...v0.1.0-alpha)

### Documentation

* Use "My Org Id" in example requests ([#83](https://github.com/m3ter-com/m3ter-sdk-go/issues/83)) ([4b53fec](https://github.com/m3ter-com/m3ter-sdk-go/commit/4b53fecd7c338c4a0707d6ee74dbebc3c59bce70))

## 0.1.0-alpha.14 (2025-04-09)

Full Changelog: [v0.1.0-alpha.13...v0.1.0-alpha.14](https://github.com/m3ter-com/m3ter-sdk-go/compare/v0.1.0-alpha.13...v0.1.0-alpha.14)

### Features

* Remove context parameter from NewClientWithServiceUserAuth ([e5b62f4](https://github.com/m3ter-com/m3ter-sdk-go/commit/e5b62f4de744a2c99e7d5249796c17c038c65bdb))
* Remove error from NewClientWithServiceUserAuth ([b9e595a](https://github.com/m3ter-com/m3ter-sdk-go/commit/b9e595a0a9cbb777839733866d66b4bf92d05a43))


### Bug Fixes

* **client:** use default params ([#81](https://github.com/m3ter-com/m3ter-sdk-go/issues/81)) ([95fa0b8](https://github.com/m3ter-com/m3ter-sdk-go/commit/95fa0b88953f2b23ec2dc2dcf3d1b47bfca17bc3))

## 0.1.0-alpha.13 (2025-04-08)

Full Changelog: [v0.1.0-alpha.12...v0.1.0-alpha.13](https://github.com/m3ter-com/m3ter-sdk-go/compare/v0.1.0-alpha.12...v0.1.0-alpha.13)

### Features

* Update README.md with service user auth ([9794f5f](https://github.com/m3ter-com/m3ter-sdk-go/commit/9794f5fede0d1fa872682b2613fd513a5b13aa38))

## 0.1.0-alpha.12 (2025-04-08)

Full Changelog: [v0.1.0-alpha.11...v0.1.0-alpha.12](https://github.com/m3ter-com/m3ter-sdk-go/compare/v0.1.0-alpha.11...v0.1.0-alpha.12)

### Features

* **api:** update contact email and package name ([#77](https://github.com/m3ter-com/m3ter-sdk-go/issues/77)) ([f9a46c8](https://github.com/m3ter-com/m3ter-sdk-go/commit/f9a46c8f3277cbb424dc736493fe7ab9f075eeda))
* **client:** support custom http clients ([#76](https://github.com/m3ter-com/m3ter-sdk-go/issues/76)) ([884f854](https://github.com/m3ter-com/m3ter-sdk-go/commit/884f854fed5b67464fe96ae44d8c6257e2687ca8))


### Bug Fixes

* **client:** return error on bad custom url instead of panic ([#75](https://github.com/m3ter-com/m3ter-sdk-go/issues/75)) ([5a8a739](https://github.com/m3ter-com/m3ter-sdk-go/commit/5a8a739e5678ad5cb79149084005e75ece1da566))
* **test:** return early after test failure ([#73](https://github.com/m3ter-com/m3ter-sdk-go/issues/73)) ([cf33469](https://github.com/m3ter-com/m3ter-sdk-go/commit/cf334697d087323b8ba5c3425ad3cb6439d12b12))


### Chores

* add request options to client tests ([#72](https://github.com/m3ter-com/m3ter-sdk-go/issues/72)) ([aa3684a](https://github.com/m3ter-com/m3ter-sdk-go/commit/aa3684a8b29e9a14bd0bf3e360b12dabc4a8f6bb))
* **docs:** improve security documentation ([#70](https://github.com/m3ter-com/m3ter-sdk-go/issues/70)) ([ccd5ec4](https://github.com/m3ter-com/m3ter-sdk-go/commit/ccd5ec43e2319514cc27ee50f1dd3c2a0a2e1575))
* fix typos ([#74](https://github.com/m3ter-com/m3ter-sdk-go/issues/74)) ([17197ef](https://github.com/m3ter-com/m3ter-sdk-go/commit/17197ef06dcbe33ee4541c67bba950fdc9c8473d))

## 0.1.0-alpha.11 (2025-03-20)

Full Changelog: [v0.1.0-alpha.10...v0.1.0-alpha.11](https://github.com/m3ter-com/m3ter-sdk-go/compare/v0.1.0-alpha.10...v0.1.0-alpha.11)

### Features

* Fix service user auth for new Stainless request option type ([4e54877](https://github.com/m3ter-com/m3ter-sdk-go/commit/4e54877cc25b25f59aeebe7934449779aac4ed64))

## 0.1.0-alpha.10 (2025-03-13)

Full Changelog: [v0.1.0-alpha.9...v0.1.0-alpha.10](https://github.com/m3ter-com/m3ter-sdk-go/compare/v0.1.0-alpha.9...v0.1.0-alpha.10)

### Features

* **client:** improve default client options support ([#63](https://github.com/m3ter-com/m3ter-sdk-go/issues/63)) ([1726176](https://github.com/m3ter-com/m3ter-sdk-go/commit/172617602f6bbcda0b8dc0685f867e8a45657f90))


### Chores

* **internal:** remove extra empty newlines ([#65](https://github.com/m3ter-com/m3ter-sdk-go/issues/65)) ([6a78f1d](https://github.com/m3ter-com/m3ter-sdk-go/commit/6a78f1dc83a76ef865035711e084cc59acb3258c))

## 0.1.0-alpha.9 (2025-03-11)

Full Changelog: [v0.1.0-alpha.8...v0.1.0-alpha.9](https://github.com/m3ter-com/m3ter-sdk-go/compare/v0.1.0-alpha.8...v0.1.0-alpha.9)

### Features

* add SKIP_BREW env var to ./scripts/bootstrap ([#58](https://github.com/m3ter-com/m3ter-sdk-go/issues/58)) ([4d65d0e](https://github.com/m3ter-com/m3ter-sdk-go/commit/4d65d0e3e4418adb750b64dbb923d8f671ac23ce))
* **client:** accept RFC6838 JSON content types ([#60](https://github.com/m3ter-com/m3ter-sdk-go/issues/60)) ([505eafe](https://github.com/m3ter-com/m3ter-sdk-go/commit/505eafe005e3833d4aacd27477ff0eff0705ae53))


### Refactors

* tidy up dependencies ([#61](https://github.com/m3ter-com/m3ter-sdk-go/issues/61)) ([605dc08](https://github.com/m3ter-com/m3ter-sdk-go/commit/605dc08c0466a0ca262bfe9089530d535967b0a7))

## 0.1.0-alpha.8 (2025-03-07)

Full Changelog: [v0.1.0-alpha.7...v0.1.0-alpha.8](https://github.com/m3ter-com/m3ter-sdk-go/compare/v0.1.0-alpha.7...v0.1.0-alpha.8)

### Features

* **client:** allow custom baseurls without trailing slash ([#55](https://github.com/m3ter-com/m3ter-sdk-go/issues/55)) ([2be2e89](https://github.com/m3ter-com/m3ter-sdk-go/commit/2be2e8944b5cc9b0ae233dd87eb347478971a6a9))

## 0.1.0-alpha.7 (2025-03-06)

Full Changelog: [v0.1.0-alpha.6...v0.1.0-alpha.7](https://github.com/m3ter-com/m3ter-sdk-go/compare/v0.1.0-alpha.6...v0.1.0-alpha.7)

### Features

* **api:** fixing warnings ([#48](https://github.com/m3ter-com/m3ter-sdk-go/issues/48)) ([0146d5e](https://github.com/m3ter-com/m3ter-sdk-go/commit/0146d5e23bad37691e7ef7ed0ff7ac184c8bf601))
* **api:** make response model names explicit ([#52](https://github.com/m3ter-com/m3ter-sdk-go/issues/52)) ([12b35c7](https://github.com/m3ter-com/m3ter-sdk-go/commit/12b35c702e365251ef0ee45e7a8d54b44339166f))
* **api:** manual updates ([#50](https://github.com/m3ter-com/m3ter-sdk-go/issues/50)) ([858e0f1](https://github.com/m3ter-com/m3ter-sdk-go/commit/858e0f1c6280aeaaac21b448d1e96d3fcf23dc80))

## 0.1.0-alpha.6 (2025-03-03)

Full Changelog: [v0.1.0-alpha.5...v0.1.0-alpha.6](https://github.com/m3ter-com/m3ter-sdk-go/compare/v0.1.0-alpha.5...v0.1.0-alpha.6)

### Features

* **api:** add missing endpoints ([#40](https://github.com/m3ter-com/m3ter-sdk-go/issues/40)) ([78a923b](https://github.com/m3ter-com/m3ter-sdk-go/commit/78a923bdcced25a367e87d7beeae624475cd7f79))
* **api:** add more endpoints ([#28](https://github.com/m3ter-com/m3ter-sdk-go/issues/28)) ([b144ab9](https://github.com/m3ter-com/m3ter-sdk-go/commit/b144ab958b845409fe9fbf11e06c3bc2884a04e6))
* **api:** add orgId path param to client settings ([#36](https://github.com/m3ter-com/m3ter-sdk-go/issues/36)) ([3a82d5d](https://github.com/m3ter-com/m3ter-sdk-go/commit/3a82d5df0fe394a7446b57296d6d1ac9829ed5e7))
* **api:** create ad hoc data export endpoint ([#34](https://github.com/m3ter-com/m3ter-sdk-go/issues/34)) ([2f4279c](https://github.com/m3ter-com/m3ter-sdk-go/commit/2f4279c00eb2b1f6641894d12da15e08951d66bc))
* **api:** snake case method names ([#41](https://github.com/m3ter-com/m3ter-sdk-go/issues/41)) ([e205a41](https://github.com/m3ter-com/m3ter-sdk-go/commit/e205a410f3c4716b2c3d14c34785e67ef94b13ec))
* **api:** Spec Update + Various Fixes ([#43](https://github.com/m3ter-com/m3ter-sdk-go/issues/43)) ([05e9595](https://github.com/m3ter-com/m3ter-sdk-go/commit/05e95955aeb7ce44e5bc0217aa21ca1d637840fa))
* **api:** update open api spec ([#39](https://github.com/m3ter-com/m3ter-sdk-go/issues/39)) ([1e18233](https://github.com/m3ter-com/m3ter-sdk-go/commit/1e18233b750652fbebb2ccc1161846fa7c913119))


### Bug Fixes

* **client:** don't truncate manually specified filenames ([#35](https://github.com/m3ter-com/m3ter-sdk-go/issues/35)) ([82c8f90](https://github.com/m3ter-com/m3ter-sdk-go/commit/82c8f904e5d073a8d6b5074ee1841dec7252ef10))
* do not call path.Base on ContentType ([#31](https://github.com/m3ter-com/m3ter-sdk-go/issues/31)) ([98b4819](https://github.com/m3ter-com/m3ter-sdk-go/commit/98b4819fbc0dc1b040c1351019fc58aea72c2637))
* fix early cancel when RequestTimeout is provided for streaming requests ([#29](https://github.com/m3ter-com/m3ter-sdk-go/issues/29)) ([f31e0c3](https://github.com/m3ter-com/m3ter-sdk-go/commit/f31e0c326a0e929e7a0e66fdeb743e98ae73c51f))


### Chores

* add UnionUnmarshaler for responses that are interfaces ([#26](https://github.com/m3ter-com/m3ter-sdk-go/issues/26)) ([ecf5f78](https://github.com/m3ter-com/m3ter-sdk-go/commit/ecf5f78f3f5e41ca1de803417ecb9bdb24642275))
* **internal:** codegen related update ([#32](https://github.com/m3ter-com/m3ter-sdk-go/issues/32)) ([d4fa04f](https://github.com/m3ter-com/m3ter-sdk-go/commit/d4fa04f3d8fd502a94dec5f0efa2fd9fa6374b0a))
* **internal:** codegen related update ([#33](https://github.com/m3ter-com/m3ter-sdk-go/issues/33)) ([dbf36b4](https://github.com/m3ter-com/m3ter-sdk-go/commit/dbf36b430a59fe37d23a0c2208988f56d20ce8aa))
* **internal:** fix devcontainers setup ([#38](https://github.com/m3ter-com/m3ter-sdk-go/issues/38)) ([7166c86](https://github.com/m3ter-com/m3ter-sdk-go/commit/7166c863766ddc20a12f6ef1830ef181f11cc5f1))
* minor change to tests ([#30](https://github.com/m3ter-com/m3ter-sdk-go/issues/30)) ([7441195](https://github.com/m3ter-com/m3ter-sdk-go/commit/744119527cb3f7e258c3e2a0879a1b4189979e19))
* org ID at the client level is required ([#45](https://github.com/m3ter-com/m3ter-sdk-go/issues/45)) ([626a859](https://github.com/m3ter-com/m3ter-sdk-go/commit/626a8599f02757dc5019a996ebc633d167bdf760))
* org ID client arg is optional ([#44](https://github.com/m3ter-com/m3ter-sdk-go/issues/44)) ([8d2850b](https://github.com/m3ter-com/m3ter-sdk-go/commit/8d2850b144f7e5d2d0de5bdcb6dd0106e30fda74))


### Documentation

* update URLs from stainlessapi.com to stainless.com ([#42](https://github.com/m3ter-com/m3ter-sdk-go/issues/42)) ([b3f9d32](https://github.com/m3ter-com/m3ter-sdk-go/commit/b3f9d32ba6bfe5ed4c5200d2edaa0d8bb8e6732c))

## 0.1.0-alpha.5 (2025-02-05)

Full Changelog: [v0.1.0-alpha.4...v0.1.0-alpha.5](https://github.com/m3ter-com/m3ter-sdk-go/compare/v0.1.0-alpha.4...v0.1.0-alpha.5)

### Features

* **api:** Update custom field type information ([#24](https://github.com/m3ter-com/m3ter-sdk-go/issues/24)) ([1e04e88](https://github.com/m3ter-com/m3ter-sdk-go/commit/1e04e88085a5cc4361f97f4518d716af7cc6bbf3))
* **client:** send `X-Stainless-Timeout` header ([#22](https://github.com/m3ter-com/m3ter-sdk-go/issues/22)) ([8a69448](https://github.com/m3ter-com/m3ter-sdk-go/commit/8a69448d87595760f5b990e4bdaf4e31294f23df))

## 0.1.0-alpha.4 (2025-02-03)

Full Changelog: [v0.1.0-alpha.3...v0.1.0-alpha.4](https://github.com/m3ter-com/m3ter-sdk-go/compare/v0.1.0-alpha.3...v0.1.0-alpha.4)

### Features

* **api:** Config update ([#19](https://github.com/m3ter-com/m3ter-sdk-go/issues/19)) ([74ca5a8](https://github.com/m3ter-com/m3ter-sdk-go/commit/74ca5a875b003bc1881ae876eee6ba56d211b946))

## 0.1.0-alpha.3 (2025-02-03)

Full Changelog: [v0.1.0-alpha.2...v0.1.0-alpha.3](https://github.com/m3ter-com/m3ter-sdk-go/compare/v0.1.0-alpha.2...v0.1.0-alpha.3)

### Features

* Add a function to automatically retrieve a service user token ([120a497](https://github.com/m3ter-com/m3ter-sdk-go/commit/120a4977194cbf04133ddfd49bcbdffc784accde))

## 0.1.0-alpha.2 (2025-01-31)

Full Changelog: [v0.1.0-alpha.1...v0.1.0-alpha.2](https://github.com/m3ter-com/m3ter-sdk-go/compare/v0.1.0-alpha.1...v0.1.0-alpha.2)

### Documentation

* document raw responses ([#14](https://github.com/m3ter-com/m3ter-sdk-go/issues/14)) ([901aec5](https://github.com/m3ter-com/m3ter-sdk-go/commit/901aec5f4b35946bdc55385ac42f837e81e7eb60))

## 0.1.0-alpha.1 (2025-01-31)

Full Changelog: [v0.0.1-alpha.1...v0.1.0-alpha.1](https://github.com/m3ter-com/m3ter-sdk-go/compare/v0.0.1-alpha.1...v0.1.0-alpha.1)

### Features

* **api:** updated OpenAPI spec ([#10](https://github.com/m3ter-com/m3ter-sdk-go/issues/10)) ([b1262dc](https://github.com/m3ter-com/m3ter-sdk-go/commit/b1262dc10d246e5e73b7ee5c97461d7259c88a64))


### Bug Fixes

* fix unicode encoding for json ([#12](https://github.com/m3ter-com/m3ter-sdk-go/issues/12)) ([a8dbab1](https://github.com/m3ter-com/m3ter-sdk-go/commit/a8dbab122865f4f686d1d03c6bc1c1c39f5b8086))

## 0.0.1-alpha.1 (2025-01-29)

Full Changelog: [v0.0.1-alpha.0...v0.0.1-alpha.1](https://github.com/m3ter-com/m3ter-sdk-go/compare/v0.0.1-alpha.0...v0.0.1-alpha.1)

### Bug Fixes

* deduplicate unknown entries in union ([#7](https://github.com/m3ter-com/m3ter-sdk-go/issues/7)) ([0600ce3](https://github.com/m3ter-com/m3ter-sdk-go/commit/0600ce3dcd16fa7e760004b42f048d8350b2da4e))
* fix apijson.Port for embedded structs ([#2](https://github.com/m3ter-com/m3ter-sdk-go/issues/2)) ([818288f](https://github.com/m3ter-com/m3ter-sdk-go/commit/818288f6f494e3dbdd6d7c870b23114a47a44943))
* fix apijson.Port for embedded structs ([#3](https://github.com/m3ter-com/m3ter-sdk-go/issues/3)) ([4515004](https://github.com/m3ter-com/m3ter-sdk-go/commit/4515004317913635d9e109e02fbf744ad25ea7ca))


### Chores

* go live ([4e94352](https://github.com/m3ter-com/m3ter-sdk-go/commit/4e943526b82389d39d40058ef012326ecf3d3a3e))
* **internal:** codegen related update ([#4](https://github.com/m3ter-com/m3ter-sdk-go/issues/4)) ([bd3c0b4](https://github.com/m3ter-com/m3ter-sdk-go/commit/bd3c0b4e999a8b2c9b4e52ce3e032798bb4fe38e))
* **internal:** codegen related update ([#6](https://github.com/m3ter-com/m3ter-sdk-go/issues/6)) ([1b06085](https://github.com/m3ter-com/m3ter-sdk-go/commit/1b0608534cece2efb97cbf655d39a89ccf75b1af))
* **tests:** skip problematic tests ([#5](https://github.com/m3ter-com/m3ter-sdk-go/issues/5)) ([ace29e2](https://github.com/m3ter-com/m3ter-sdk-go/commit/ace29e23d457419fed994abce2519286126d236f))
