package functions

import "fmt"

// This is an example. Replace thie following with whatever steps are needed to
// install required components into
// const dockerfileLines = `RUN apt-get update && \
// apt-get install gnupg apt-transport-https lsb-release software-properties-common -y && \
// echo "deb [arch=amd64] https://packages.microsoft.com/repos/azure-cli/ stretch main" | \
//    tee /etc/apt/sources.list.d/azure-cli.list && \
// apt-key --keyring /etc/apt/trusted.gpg.d/Microsoft.gpg adv \
// 	--keyserver packages.microsoft.com \
// 	--recv-keys BC528686B50D79E339D3721CEB3E94ADBE1229CF && \
// apt-get update && apt-get install azure-cli
// `

// Build will generate the necessary Dockerfile lines
// for an invocation image using this mixin
func (m *Mixin) Build() error {

	fmt.Fprintln(m.Out, `RUN apt-get update`)
	fmt.Fprintln(m.Out, `RUN apt-get install -y curl gnupg gnupg2 apt-transport-https`)
	//fmt.Fprintln(m.Out, `RUN apt-get install -y `)
	fmt.Fprintln(m.Out, `RUN curl https://packages.microsoft.com/keys/microsoft.asc -o foo.asc`)
	fmt.Fprintln(m.Out, `RUN gpg foo.asc > microsoft.gpg`)
	fmt.Fprintln(m.Out, `RUN mv microsoft.gpg /etc/apt/trusted.gpg.d/microsoft.gpg`)
	fmt.Fprintln(m.Out, `RUN echo "deb [arch=amd64] https://packages.microsoft.com/debian/9/prod stretch main" > /etc/apt/sources.list.d/dotnetdev.list`)
	//fmt.Fprintln(m.Out, `RUN cat /etc/apt/sources.list.d`)
	fmt.Fprintln(m.Out, `RUN apt-get update`)
	fmt.Fprintln(m.Out, `RUN apt-get install -y --allow-unauthenticated azure-functions-core-tools`)

	return nil
}
