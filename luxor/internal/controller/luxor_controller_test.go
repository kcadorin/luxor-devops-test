package controller

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client" // Import client package
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	corev1 "k8s.io/api/core/v1" // Import corev1 package
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	luxorv1 "luxor.test/test/api/v1"
)

var _ = Describe("Luxor Controller", func() {
	Context("When reconciling a resource", func() {
		const resourceName = "test-luxor"

		ctx := context.Background()

		typeNamespacedName := types.NamespacedName{
			Name:      resourceName,
			Namespace: "default", // Modify as needed
		}
		luxor := &luxorv1.Luxor{}

		BeforeEach(func() {
			By("creating the custom resource for the Luxor")
			err := k8sClient.Get(ctx, typeNamespacedName, luxor)
			if err != nil && errors.IsNotFound(err) {
				resource := &luxorv1.Luxor{
					ObjectMeta: metav1.ObjectMeta{
						Name:      resourceName,
						Namespace: "default",
					},
					Spec: luxorv1.LuxorSpec{
						Replicas: 3, // Modify as needed
						Config: luxorv1.HTTPServerConfig{
							Hosts: []luxorv1.HostConfig{
								{
									Host: "example.com",
									Port: 80,
								},
								// Add other hosts as needed
							},
						},
					},
				}
				Expect(k8sClient.Create(ctx, resource)).To(Succeed())
			}
		})

		AfterEach(func() {
			By("cleaning up the resources after each test")
			resource := &luxorv1.Luxor{}
			err := k8sClient.Get(ctx, typeNamespacedName, resource)
			Expect(err).NotTo(HaveOccurred())

			Expect(k8sClient.Delete(ctx, resource)).To(Succeed())
		})

		It("should successfully reconcile the resource", func() {
			By("reconciling the created resource")
			controllerReconciler := &LuxorReconciler{
				Client: k8sClient,
				Scheme: k8sClient.Scheme(),
			}

			_, err := controllerReconciler.Reconcile(ctx, reconcile.Request{
				NamespacedName: typeNamespacedName,
			})
			Expect(err).NotTo(HaveOccurred())

			By("verifying the resources created by the controller")

			// Check if pods are created
			podList := &corev1.PodList{}
			err = k8sClient.List(ctx, podList, client.InNamespace("default"), client.MatchingLabels{"app": resourceName})
			Expect(err).NotTo(HaveOccurred())
			Expect(len(podList.Items)).To(Equal(3)) // Expect 3 pods to be created

			// Add more specific assertions as needed
		})
	})
})
