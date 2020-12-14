package models

import (
	"reflect"
	"testing"
	"time"
)

func TestNewAnnotation(t *testing.T) {
	tests := []struct {
		name string
		want ESAnnotationBuilder
	}{
		{
			name: "Should return a new annotation builder",
			want: &esAnnotationBuilder{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAnnotation(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAnnotation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_esAnnotationBuilder_WithText(t *testing.T) {
	type fields struct {
		text        string
		tags        []string
		environment string
		role        string
		region      string
		provider    string
		hostname    string
	}
	type args struct {
		text string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   ESAnnotationBuilder
	}{
		{
			name:   "Should add the porper test",
			fields: fields{},
			args: args{
				text: "bladibla",
			},
			want: &esAnnotationBuilder{text: "bladibla"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &esAnnotationBuilder{
				text:        tt.fields.text,
				tags:        tt.fields.tags,
				environment: tt.fields.environment,
				role:        tt.fields.role,
				region:      tt.fields.region,
				provider:    tt.fields.provider,
				hostname:    tt.fields.hostname,
			}
			if got := e.WithText(tt.args.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_esAnnotationBuilder_WithTag(t *testing.T) {
	type fields struct {
		text        string
		tags        []string
		environment string
		role        string
		region      string
		provider    string
		hostname    string
	}
	type args struct {
		tag string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   ESAnnotationBuilder
	}{
		{
			name: "Should add the tag to the current list",
			fields: fields{
				tags: []string{"bladibla"},
			},
			args: args{"duh"},
			want: &esAnnotationBuilder{tags: []string{"bladibla", "duh"}},
		},
		{
			name:   "Should create a list when empty",
			fields: fields{},
			args:   args{"duh"},
			want:   &esAnnotationBuilder{tags: []string{"duh"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &esAnnotationBuilder{
				text:        tt.fields.text,
				tags:        tt.fields.tags,
				environment: tt.fields.environment,
				role:        tt.fields.role,
				region:      tt.fields.region,
				provider:    tt.fields.provider,
				hostname:    tt.fields.hostname,
			}
			if got := e.WithTag(tt.args.tag); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithTag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_esAnnotationBuilder_WithTags(t *testing.T) {
	type fields struct {
		text        string
		tags        []string
		environment string
		role        string
		region      string
		provider    string
		hostname    string
	}
	type args struct {
		tags []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   ESAnnotationBuilder
	}{
		{
			name: "Should add the proper list of tags to the current tags",
			fields: fields{
				tags: []string{"bladibla"},
			},
			args: args{[]string{"duh"}},
			want: &esAnnotationBuilder{tags: []string{"bladibla", "duh"}},
		},
		{
			name:   "Should add the list is empty",
			fields: fields{},
			args:   args{[]string{"duh"}},
			want:   &esAnnotationBuilder{tags: []string{"duh"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &esAnnotationBuilder{
				text:        tt.fields.text,
				tags:        tt.fields.tags,
				environment: tt.fields.environment,
				role:        tt.fields.role,
				region:      tt.fields.region,
				provider:    tt.fields.provider,
				hostname:    tt.fields.hostname,
			}
			if got := e.WithTags(tt.args.tags); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithTags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_esAnnotationBuilder_WithEnvironment(t *testing.T) {
	type fields struct {
		text        string
		tags        []string
		environment string
		role        string
		region      string
		provider    string
		hostname    string
	}
	type args struct {
		env string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   ESAnnotationBuilder
	}{
		{
			name:   "should add the proper env value",
			fields: fields{},
			args:   args{env: "dev"},
			want:   &esAnnotationBuilder{environment: "dev"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &esAnnotationBuilder{
				text:        tt.fields.text,
				tags:        tt.fields.tags,
				environment: tt.fields.environment,
				role:        tt.fields.role,
				region:      tt.fields.region,
				provider:    tt.fields.provider,
				hostname:    tt.fields.hostname,
			}
			if got := e.WithEnvironment(tt.args.env); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithEnvironment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_esAnnotationBuilder_WithHostname(t *testing.T) {
	type fields struct {
		text        string
		tags        []string
		environment string
		role        string
		region      string
		provider    string
		hostname    string
	}
	type args struct {
		hostname string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   ESAnnotationBuilder
	}{
		{
			name:   "should add the proper hostname value",
			fields: fields{},
			args:   args{hostname: "bladibla"},
			want:   &esAnnotationBuilder{hostname: "bladibla"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &esAnnotationBuilder{
				text:        tt.fields.text,
				tags:        tt.fields.tags,
				environment: tt.fields.environment,
				role:        tt.fields.role,
				region:      tt.fields.region,
				provider:    tt.fields.provider,
				hostname:    tt.fields.hostname,
			}
			if got := e.WithHostname(tt.args.hostname); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithHostname() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_esAnnotationBuilder_WithProvider(t *testing.T) {
	type fields struct {
		text        string
		tags        []string
		environment string
		role        string
		region      string
		provider    string
		hostname    string
	}
	type args struct {
		provider string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   ESAnnotationBuilder
	}{
		{
			name:   "should add the proper provider value",
			fields: fields{},
			args:   args{provider: "bladibla"},
			want:   &esAnnotationBuilder{provider: "bladibla"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &esAnnotationBuilder{
				text:        tt.fields.text,
				tags:        tt.fields.tags,
				environment: tt.fields.environment,
				role:        tt.fields.role,
				region:      tt.fields.region,
				provider:    tt.fields.provider,
				hostname:    tt.fields.hostname,
			}
			if got := e.WithProvider(tt.args.provider); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithProvider() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_esAnnotationBuilder_WithRegion(t *testing.T) {
	type fields struct {
		text        string
		tags        []string
		environment string
		role        string
		region      string
		provider    string
		hostname    string
	}
	type args struct {
		region string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   ESAnnotationBuilder
	}{
		{
			name:   "should add the proper region value",
			fields: fields{},
			args:   args{region: "bladibla"},
			want:   &esAnnotationBuilder{region: "bladibla"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &esAnnotationBuilder{
				text:        tt.fields.text,
				tags:        tt.fields.tags,
				environment: tt.fields.environment,
				role:        tt.fields.role,
				region:      tt.fields.region,
				provider:    tt.fields.provider,
				hostname:    tt.fields.hostname,
			}
			if got := e.WithRegion(tt.args.region); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithRegion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_esAnnotationBuilder_WithRole(t *testing.T) {
	type fields struct {
		text        string
		tags        []string
		environment string
		role        string
		region      string
		provider    string
		hostname    string
	}
	type args struct {
		role string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   ESAnnotationBuilder
	}{
		{
			name:   "should add the proper role value",
			fields: fields{},
			args:   args{role: "bladibla"},
			want:   &esAnnotationBuilder{role: "bladibla"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &esAnnotationBuilder{
				text:        tt.fields.text,
				tags:        tt.fields.tags,
				environment: tt.fields.environment,
				role:        tt.fields.role,
				region:      tt.fields.region,
				provider:    tt.fields.provider,
				hostname:    tt.fields.hostname,
			}
			if got := e.WithRole(tt.args.role); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithRole() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_esAnnotationBuilder_Build(t *testing.T) {
	now := time.Now()
	type fields struct {
		text        string
		tags        []string
		environment string
		role        string
		region      string
		provider    string
		hostname    string
	}
	tests := []struct {
		name   string
		fields fields
		want   *EsAnnotation
	}{
		{
			name: "Should build the annotation",
			fields: fields{
				text:        "sample text",
				tags:        []string{"bladibla", "duh"},
				environment: "dev",
				role:        "test",
				region:      "eu-west-1",
				provider:    "aws",
				hostname:    "localdesktop",
			},
			want: &EsAnnotation{
				Time:        &now,
				Text:        "sample text",
				Tags:        []string{"bladibla", "duh"},
				Environment: "dev",
				Role:        "test",
				Region:      "eu-west-1",
				Provider:    "aws",
				Hostname:    "localdesktop",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &esAnnotationBuilder{
				text:        tt.fields.text,
				tags:        tt.fields.tags,
				environment: tt.fields.environment,
				role:        tt.fields.role,
				region:      tt.fields.region,
				provider:    tt.fields.provider,
				hostname:    tt.fields.hostname,
			}
			annotation := e.Build()
			if got := annotation; !reflect.DeepEqual(got, tt.want) {
				if now.Format("2006-01-02 15:03:04") != annotation.Time.Format("2006-01-02 15:03:04") {
					t.Errorf("Build() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestEsAnnotation_String(t *testing.T) {
	t1, _ := time.Parse("2006-01-02", "2020-12-17")
	type fields struct {
		Time        *time.Time
		Text        string
		Tags        []string
		Environment string
		Role        string
		Region      string
		Provider    string
		Hostname    string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Should return a json string representation of the struct",
			fields: fields{
				Time:        &t1,
				Text:        "some text",
				Tags:        []string{"bladibla"},
				Environment: "dev",
				Role:        "testrole",
				Region:      "eu-west-1",
				Provider:    "aws",
				Hostname:    "bladibla",
			},
			want: `{"time":"2020-12-17T00:00:00Z","text":"some text","tags":["bladibla"],"environment":"dev","role":"testrole","region":"eu-west-1","provider":"aws","hostname":"bladibla"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &EsAnnotation{
				Time:        tt.fields.Time,
				Text:        tt.fields.Text,
				Tags:        tt.fields.Tags,
				Environment: tt.fields.Environment,
				Role:        tt.fields.Role,
				Region:      tt.fields.Region,
				Provider:    tt.fields.Provider,
				Hostname:    tt.fields.Hostname,
			}
			if got := e.String(); got != tt.want {
				t.Errorf("String() = %v, \n want %v", got, tt.want)
			}
		})
	}
}

func TestEsAnnotation_GetIndexName(t *testing.T) {
	t1, _ := time.Parse("2006-01-02", "2020-12-17")
	type fields struct {
		Time        *time.Time
		Text        string
		Tags        []string
		Environment string
		Role        string
		Region      string
		Provider    string
		Hostname    string
	}
	type args struct {
		indexPattern string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "Should return the proper index name",
			fields: fields{
				Time: &t1,
			},
			args: args{
				indexPattern: "bladibla-2006.01",
			},
			want: "bladibla-2020.12",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &EsAnnotation{
				Time:        tt.fields.Time,
				Text:        tt.fields.Text,
				Tags:        tt.fields.Tags,
				Environment: tt.fields.Environment,
				Role:        tt.fields.Role,
				Region:      tt.fields.Region,
				Provider:    tt.fields.Provider,
				Hostname:    tt.fields.Hostname,
			}
			if got := e.GetIndexName(tt.args.indexPattern); got != tt.want {
				t.Errorf("GetIndexName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEsAnnotation_getDocumentID(t *testing.T) {
	t1, _ := time.Parse("2006-01-02", "2020-12-17")
	type fields struct {
		Time        *time.Time
		Text        string
		Tags        []string
		Environment string
		Role        string
		Region      string
		Provider    string
		Hostname    string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Should compute the document ID",
			fields: fields{
				Time: &t1,
			},
			want: "3c2016306aea2fb7eccec038012ed2d5",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &EsAnnotation{
				Time:        tt.fields.Time,
				Text:        tt.fields.Text,
				Tags:        tt.fields.Tags,
				Environment: tt.fields.Environment,
				Role:        tt.fields.Role,
				Region:      tt.fields.Region,
				Provider:    tt.fields.Provider,
				Hostname:    tt.fields.Hostname,
			}
			if got := e.getDocumentID(); got != tt.want {
				t.Errorf("getDocumentID() = %v, want %v", got, tt.want)
			}
		})
	}
}
